$(function() {
    $('#loginForm').on('submit', function(event) {
        event.preventDefault();

        const identifier = $('#identifier').val().trim();
        const password = $('#password').val();

        if (!identifier || !password) {
            Swal.fire("Oops...", "Please fill in all fields!", "error");
            return;
        }

        $.ajax({
            url: '/login',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({ identifier, password }),
            success: function() {
                window.location.href = '/feed';
            },
            error: function(xhr) {
                let msg = "Login failed. Please check your credentials.";
                if (xhr.responseJSON && xhr.responseJSON.error) {
                    msg = xhr.responseJSON.error;
                }
                Swal.fire("Oops...", msg, "error");
            }
        });
    });
});