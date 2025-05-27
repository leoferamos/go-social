$(function() {
    $("#registerForm").on("submit", function(e) {
        e.preventDefault();

        var name = $("#name").val().trim();
        var username = $("#username").val().trim();
        var email = $("#email").val().trim();
        var password = $("#password").val();
        var confirmPassword = $("#confirmPassword").val();

        if (!name || !username || !email || !password || !confirmPassword) {
            alert("Please fill in all fields.");
            return;
        }

        if (password !== confirmPassword) {
            $("#passwordError").show();
            $("#confirmPassword").focus();
            return;
        } else {
            $("#passwordError").hide();
        }

        var userData = {
            name: name,
            username: username,
            email: email,
            password: password
        };

        $.ajax({
            url: "/register",
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify(userData),
            success: function(response) {
                window.location.href = "/login";
            },
            error: function(xhr) {
                let msg = "Registration failed. Please check your data.";
                if (xhr.responseJSON && xhr.responseJSON.error) {
                    msg = xhr.responseJSON.error;
                }
                alert(msg);
            }
        });
    });

    $("#confirmPassword, #password").on("input", function() {
        $("#passwordError").hide();
    });
});