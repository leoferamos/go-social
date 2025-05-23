$('#login').on('submit', function(event) {
    event.preventDefault();

    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            identifier: $('#identifier').val(),
            password: $('#password').val()
        }
    }).done(function() {
        window.location.href = '/home';
    }).fail(function(xhr) {
        alert('Login failed: ' + xhr.responseText);
    });
});