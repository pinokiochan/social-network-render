document.getElementById('emailForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    const users = formData.get('users').split(',').map(email => email.trim());
    formData.delete('users');
    users.forEach(user => formData.append('users[]', user));

    // Get the JWT token from localStorage
    const token = localStorage.getItem('token');  // Make sure you have stored the token previously

    if (!token) {
        alert('No token found, please login again.');
        return;
    }

    try {
        const response = await fetch('https://social-network-2.onrender.com/api/admin/broadcast-to-selected', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,  // Use thdee JWT token from localStorage
            },
            body: formData,
        });

        const result = await response.json();

        if (response.ok) {
            alert('Emails sent successfully!');
        } else {
            // Handle the error response
            alert(`Failed to send emails: ${result.message}`);
        }
    } catch (err) {
        // Catch and log any errors in the request or response
        console.error('Error sending emails:', err);
        alert('An error occurred. Please try again.');
    }
});
