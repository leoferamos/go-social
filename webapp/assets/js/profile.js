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

$(function() {

    $('#edit-profile-btn').on('click', function() {
        const userId = $(this).data('user-id');
        const defaultAvatar = "{{ .Profile.User.AvatarURL }}";
        const defaultBanner = "{{ .Profile.User.BannerURL }}";

        $.get(`/users/${userId}`, function(user) {
            $('#edit-name').val(user.name);
            $('#edit-username').val(user.username);
            $('#edit-email').val(user.email);
            $('#edit-bio').val(user.bio || '');

            $('#edit-profile-avatar').attr('src', user.avatar_url || defaultAvatar);
            $('#edit-profile-banner').attr('src', user.banner_url || defaultBanner);

            $('#edit-profile-overlay, #edit-profile-modal').fadeIn(150); 
        });
    });


    $('#edit-profile-overlay, #close-edit-profile').on('click', function() {
        $('#edit-profile-overlay, #edit-profile-modal').fadeOut(150); 
    });

    $('#edit-profile-form').on('submit', function(e) {
        e.preventDefault();
        const userId = $('#edit-profile-btn').data('user-id');
        const oldUsername = "{{ index .Profile.User \"username\" }}";
        const newUsername = $('#edit-username').val().trim();

        const data = {
            name: $('#edit-name').val().trim(),
            username: newUsername,
            email: $('#edit-email').val().trim(),
            bio: $('#edit-bio').val().trim()
        };
        $.ajax({
            url: `/users/${userId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function() {
                if (oldUsername !== newUsername) {
                    window.location.href = `/profile/${newUsername}`;
                } else {
                    location.reload();
                }
            },
            error: function(xhr) {
                let msg = 'Error updating profile.';
                if (xhr.responseJSON && xhr.responseJSON.error) {
                    if (xhr.responseJSON.error.includes('Duplicate entry') && xhr.responseJSON.error.includes('username')) {
                        msg = 'This username is already taken. Please choose another one.';
                    } else {
                        msg = xhr.responseJSON.error;
                    }
                }
                Swal.fire("Oops...", msg, "error");
            }
        });
    });

    $('.show-followers').on('click', function(e) {
        e.preventDefault();
        const userId = $(this).data('user-id');
        $('#followers-modal-title').text('Followers');
        $('#followers-modal-content').html('<div class="text-center my-4"><div class="spinner-border"></div></div>');
        $('#followers-modal-overlay, #followers-modal').fadeIn(150);

        $.get(`/users/${userId}/followers`, function(followers) {
            renderFollowersList(followers);
        }).fail(function() {
            $('#followers-modal-content').html('<div class="text-danger text-center my-4">Failed to load followers.</div>');
        });
    });


    $('.show-following').on('click', function(e) {
        e.preventDefault();
        const userId = $(this).data('user-id');
        $('#followers-modal-title').text('Following');
        $('#followers-modal-content').html('<div class="text-center my-4"><div class="spinner-border"></div></div>');
        $('#followers-modal-overlay, #followers-modal').fadeIn(150);

        $.get(`/users/${userId}/following`, function(following) {
            renderFollowersList(following);
        }).fail(function() {
            $('#followers-modal-content').html('<div class="text-danger text-center my-4">Failed to load following.</div>');
        });
    });

    $('#followers-modal-overlay, #close-followers-modal').on('click', function() {
        $('#followers-modal-overlay, #followers-modal').fadeOut(150);
    });

    function renderFollowersList(users) {
        if (!users.length) {
            $('#followers-modal-content').html('<div class="text-center my-4 text-muted">No users found.</div>');
            return;
        }
        let html = '<ul class="list-group list-group-flush">';
        users.forEach(function(user) {
            html += `
                <li class="list-group-item d-flex align-items-center gap-3">
                    <img src="${user.avatar_url || user.AvatarURL || '/assets/img/avatar-placeholder.png'}" alt="Avatar" class="rounded-circle" style="width:40px;height:40px;object-fit:cover;">
                    <div>
                        <div><b>${user.name || user.Name}</b></div>
                        <div class="text-muted small">@${user.username || user.Username}</div>
                    </div>
                </li>
            `;
        });
        html += '</ul>';
        $('#followers-modal-content').html(html);
    }
}); 