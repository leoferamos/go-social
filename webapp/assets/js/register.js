$(function() {
    $("#registerForm").on("submit", function(e) {
        e.preventDefault();

        var password = $("#password").val();
        var confirmPassword = $("#confirm_password").val();

        if (password !== confirmPassword) {
            $("#passwordError").show();
            $("#confirm_password").focus();
            return;
        } else {
            $("#passwordError").hide();
        }

        var userData = {
            name: $("#name").val(),
            username: $("#username").val(),
            email: $("#email").val(),
            password: $("#password").val()
        };

        $.ajax({
            url: "/register",
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify(userData),
            success: function(response) {
                alert("Registration successful!");
                window.location.href = "/login";
            },
            error: function(xhr) {
                alert("Registration failed: " + xhr.responseText);
            }
        });
    });

    $("#confirm_password, #password").on("input", function() {
        $("#passwordError").hide();
    });
});