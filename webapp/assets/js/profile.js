$(function () {
    const $editProfileModal = $('#edit-profile-modal');
    const $editProfileOverlay = $('#edit-profile-overlay');
    const $followersModal = $('#followers-modal');
    const $followersOverlay = $('#followers-modal-overlay');
    const $followersModalTitle = $('#followers-modal-title');
    const $followersModalContent = $('#followers-modal-content');
    const $editProfileBtn = $('#edit-profile-btn');
    const $editProfileForm = $('#edit-profile-form');
    const $followBtn = $('#follow-btn');

    // FOLLOW/UNFOLLOW BUTTON
    if ($followBtn.length) {
        updateFollowBtn($followBtn, Boolean($followBtn.data('is-following')));

        $followBtn.on('click', function () {
            const userId = $followBtn.data('user-id');
            const isFollowing = $followBtn.hasClass('following');

            $.ajax({
                url: isFollowing ? `/users/${userId}/unfollow` : `/users/${userId}/follow`,
                method: 'POST',
                success: function (data) {
                    updateFollowBtn($followBtn, data.following);
                },
                error: function (xhr) {
                    showError(xhr, 'An error occurred while trying to follow/unfollow.');
                }
            });
        });
    }

    function updateFollowBtn($btn, isFollowing) {
        if (isFollowing) {
            $btn.addClass('following btn-outline-primary')
                .removeClass('btn-primary')
                .text('Unfollow');
        } else {
            $btn.removeClass('following btn-outline-primary')
                .addClass('btn-primary')
                .text('Follow');
        }
    }

    // EDIT PROFILE MODAL
    $editProfileBtn.on('click', function () {
        const userId = $(this).data('user-id');
        const defaultAvatar = $editProfileBtn.data('avatar-url') || '/assets/img/avatar-placeholder.png';
        const defaultBanner = $editProfileBtn.data('banner-url') || '/assets/img/banner-placeholder.png';

        $.get(`/users/${userId}`, function (user) {
            $('#edit-name').val(user.name);
            $('#edit-username').val(user.username);
            $('#edit-email').val(user.email);
            $('#edit-bio').val(user.bio || '');
            $('#edit-profile-avatar').attr('src', user.avatar_url || defaultAvatar);
            $('#edit-profile-banner').attr('src', user.banner_url || defaultBanner);
            $editProfileOverlay.add($editProfileModal).fadeIn(150);
        });
    });

    $editProfileOverlay.add('#close-edit-profile').on('click', function () {
        $editProfileOverlay.add($editProfileModal).fadeOut(150);
    });

    $editProfileForm.on('submit', function (e) {
        e.preventDefault();
        const userId = $editProfileBtn.data('user-id');
        const oldUsername = $editProfileBtn.data('username');
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
            success: function () {
                if (oldUsername !== newUsername) {
                    window.location.href = `/profile/${newUsername}`;
                } else {
                    location.reload();
                }
            },
            error: function (xhr) {
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

    // FOLLOWERS / FOLLOWING MODALS
    $(document).on('click', '.show-followers', function (e) {
        e.preventDefault();
        const userId = $(this).data('user-id');
        $followersModalTitle.text('Followers');
        showModalLoading();
        $followersOverlay.add($followersModal).fadeIn(150);

        $.get(`/users/${userId}/followers`, function(users) {
            renderFollowersList(users, "This user has no followers yet.");
        }).fail(function () {
            showModalError('Failed to load followers.');
        });
    });

    $(document).on('click', '.show-following', function (e) {
        e.preventDefault();
        const userId = $(this).data('user-id');
        $followersModalTitle.text('Following');
        showModalLoading();
        $followersOverlay.add($followersModal).fadeIn(150);

        $.get(`/users/${userId}/following`, function(users) {
            renderFollowersList(users, "This user is not following anyone yet.");
        }).fail(function () {
            showModalError('Failed to load following.');
        });
    });

    $followersOverlay.add('#close-followers-modal').on('click', function () {
        $followersOverlay.add($followersModal).fadeOut(150);
    });

    // UTILS
    function renderFollowersList(users, emptyMsg) {
        if (!Array.isArray(users) || !users.length) {
            $followersModalContent.html(`<div class="text-center my-4" style="color: #E7E9EA;">${emptyMsg}</div>`);
            return;
        }
        const html = users.map(user => `
            <li class="list-group-item d-flex align-items-center gap-3">
                <img src="${user.avatar_url || user.AvatarURL || '/assets/img/avatar-placeholder.png'}" alt="Avatar" class="rounded-circle" style="width:40px;height:40px;object-fit:cover;">
                <div>
                    <div><b>${user.name || user.Name}</b></div>
                    <div class="text-muted small">@${user.username || user.Username}</div>
                </div>
            </li>
        `).join('');
        $followersModalContent.html(`<ul class="list-group list-group-flush">${html}</ul>`);
    }

    function showModalLoading() {
        $followersModalContent.html('<div class="text-center my-4"><div class="spinner-border"></div></div>');
    }

    function showModalError(msg) {
        $followersModalContent.html(`<div class="text-danger text-center my-4">${msg}</div>`);
    }

    function showError(xhr, defaultMsg) {
        let msg = defaultMsg;
        if (xhr.responseJSON && xhr.responseJSON.error) {
            msg = xhr.responseJSON.error;
        }
        Swal.fire("Oops...", msg, "error");
    }
});
