$(document).ready(function() {
    const $btn = $('#follow-btn');
    if ($btn.length) {
        const isFollowing = Boolean($btn.data('is-following'));
        if (isFollowing) {
            $btn
                .addClass('following btn-outline-primary')
                .removeClass('btn-primary')
                .text('Unfollow');
        } else {
            $btn
                .removeClass('following btn-outline-primary')
                .addClass('btn-primary')
                .text('Follow');
        }
    }
    $btn.on('click', follow);
});

function follow() {
    const $btn = $(this);
    const userId = $btn.data('user-id');
    const isFollowing = $btn.hasClass('following');

    $.ajax({
        url: isFollowing ? `/users/${userId}/unfollow` : `/users/${userId}/follow`,
        method: 'POST',
        success: function(data) {
            if (data.following) {
                $btn
                    .addClass('following btn-outline-primary')
                    .removeClass('btn-primary')
                    .text('Unfollow');
            } else {
                $btn
                    .removeClass('following btn-outline-primary')
                    .addClass('btn-primary')
                    .text('Follow');
            }
        },
        error: function(xhr) {
            let msg = 'An error occurred while trying to follow/unfollow.';
            if (xhr.responseJSON && xhr.responseJSON.error) {
                msg = xhr.responseJSON.error;
            }
            Swal.fire("Oops...", msg, "error");
        }
    });
}