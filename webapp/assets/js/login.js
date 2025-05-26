$('#login').on('submit', function(event) {
    event.preventDefault();

    $.ajax({
        url: '/login',
        method: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({
            identifier: $('#identifier').val(),
            password: $('#password').val()
        })
    }).done(function() {
        window.location.href = '/feed';
    }).fail(function(xhr) {
        alert('Login failed: ' + xhr.responseText);
    });
});