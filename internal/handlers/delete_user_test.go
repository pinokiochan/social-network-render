package handlers_test

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pinokiochan/social-network-render/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	// Создание mock-базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Не удалось создать mock-базу данных: %s", err)
	}
	defer db.Close()

	// Моделируем выполнение SQL запроса на удаление пользователя
	mock.ExpectExec("DELETE FROM users WHERE id = \\$1").
		WithArgs(123).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Моделируем успешное удаление (1 строка затронута)

	// Создание обработчика
	handler := handlers.NewAdminHandler(db, nil)

	// Создание запроса DELETE
	req, err := http.NewRequest("DELETE", "/api/admin/users/delete?id=123", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %s", err)
	}

	// Регистрируем ответ
	rec := httptest.NewRecorder()

	// Вызов обработчика
	handler.DeleteUser(rec, req)

	// Проверка статуса ответа (должен быть 200 OK)
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидался статус 200 OK, получен: %d", rec.Code)
	}

	// Проверка содержимого ответа
	expected := map[string]string{"message": "User deleted successfully"}
	var actual map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Не удалось разобрать ответ: %s", err)
	}
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("Ожидался ответ %v, получен %v", expected, actual)
	}

	// Проверка всех ожиданий mock-базы данных
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Не все ожидания были выполнены: %s", err)
	}
}
