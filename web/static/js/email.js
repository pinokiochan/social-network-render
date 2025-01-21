import AdminAuth from './admin.js'; // Укажите правильный путь к файлу

document.getElementById('emailForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    const users = formData.get('users').split(',').map(email => email.trim());
    formData.delete('users');
    users.forEach(user => formData.append('users[]', user));

    // Получаем токен через AdminAuth
    const token = AdminAuth.getToken();
    if (!token) {
        alert('Токен не найден! Убедитесь, что вы вошли в систему.');
        return;
    }

    try {
        const response = await fetch('https://social-network-2.onrender.com/api/admin/broadcast-to-selected', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`, // Передаём токен в заголовок
            },
            body: formData,
        });

        const result = await response.json();

        if (response.ok) {
            alert('Emails sent successfully!');
        } else {
            alert(`Failed to send emails: ${result.message}`);
        }
    } catch (err) {
        console.error('Error sending emails:', err);
        alert('An error occurred. Please try again.');
    }
});
