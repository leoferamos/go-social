$(function() {
    $("#registerForm").on("submit", function(e) {
        var password = $("#password").val();
        var confirmPassword = $("#confirm_password").val();

        if (password !== confirmPassword) {
            $("#passwordError").show();
            $("#confirm_password").focus();
            e.preventDefault();
        } else {
            $("#passwordError").hide();
        }
    });


    $("#confirm_password, #password").on("input", function() {
        $("#passwordError").hide();
    });

    $.ajax({
        url: "register",
        method: "POST",
        data: {
            name: $("#name").val(),
            username: $("#username").val(),
            email: $("#email").val(),
            password: $("#password").val()
        },
        success: function(response) {
            // Handle successful registration
        },
        error: function(xhr, status, error) {
            // Handle registration error
        }
    });
});