$(`#login`).on('submit', login);

function login(event) {
    event.preventDefault();
    
    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            username: $('#username').val(),
            password: $('#password').val()
        }
    }).done(function() {
        window.location.href = '/home';
    }
    ).fail(function() {
        alert('Login failed. Please check your username and password.');
    });
}