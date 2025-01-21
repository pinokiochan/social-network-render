document.getElementById('emailForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    const users = formData.get('users').split(',').map(email => email.trim());
    formData.delete('users');
    users.forEach(user => formData.append('users[]', user));

    // Получаем токен через AdminAuth
    const token = AdminAuth.getToken();

    if (!token) {
        console.error('Token not found');
        alert('You are not authorized. Please log in first.');
        return;
    }

    try {
        const response = await fetch('https://social-network-2.onrender.com/api/admin/broadcast-to-selected', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`, // Используем токен из AdminAuth
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
