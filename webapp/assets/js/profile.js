$(function () {
    // --- Cached selectors ---
    const $editProfileModal = $('#edit-profile-modal');
    const $editProfileOverlay = $('#edit-profile-overlay');
    const $followersModal = $('#followers-modal');
    const $followersOverlay = $('#followers-modal-overlay');
    const $followersModalTitle = $('#followers-modal-title');
    const $followersModalContent = $('#followers-modal-content');
    const $editProfileBtn = $('#edit-profile-btn');
    const $editProfileForm = $('#edit-profile-form');
    const $followBtn = $('#follow-btn');

    const currentUserId = Number($('body').data('current-user-id'));

    // --- Follow/Unfollow Button on Profile ---
    if ($followBtn.length) {
        updateFollowBtn($followBtn, Boolean($followBtn.data('is-following')));
        $followBtn.on('click', function () {
            const userId = $followBtn.data('user-id');
            const isFollowing = $followBtn.hasClass('following');
            toggleFollow($followBtn, userId, isFollowing, updateFollowBtn);
        });
    }

    // --- Edit Profile Modal ---
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

    // --- Followers / Following Modals ---
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

    // --- Follow/Unfollow Button in List ---
    $(document).on('click', '.follow-list-btn', function () {
        const $btn = $(this);
        const userId = $btn.data('user-id');
        const isFollowing = $btn.hasClass('following');
        toggleFollow($btn, userId, isFollowing, updateFollowListBtn);
    });

    // --- Helpers ---
    function toggleFollow($btn, userId, isFollowing, updateBtnFn) {
        $.ajax({
            url: isFollowing ? `/users/${userId}/unfollow` : `/users/${userId}/follow`,
            method: 'POST',
            success: function (data) {
                updateBtnFn($btn, data.following ?? data.is_following);
            },
            error: function (xhr) {
                showError(xhr, 'An error occurred while trying to follow/unfollow.');
            }
        });
    }

    function updateFollowBtn($btn, isFollowing) {
        if (isFollowing) {
            $btn.addClass('following btn-outline-primary')
                .removeClass('btn-primary')
                .text('Unfollow')
                .data('is-following', true);
        } else {
            $btn.removeClass('following btn-outline-primary')
                .addClass('btn-primary')
                .text('Follow')
                .data('is-following', false);
        }
    }

    function updateFollowListBtn($btn, isFollowing) {
        updateFollowBtn($btn, isFollowing);
    }

    function userListItem(user, isFollowing) {
        const username = user.username || user.Username;
        const name = user.name || user.Name;
        const avatar = user.avatar_url || user.AvatarURL || '/assets/img/avatar-placeholder.png';
        const isSelf = currentUserId === user.id;
        return `
            <li class="list-group-item d-flex align-items-center gap-3 justify-content-between" data-user-id="${user.id}">
                <div class="d-flex align-items-center gap-3">
                    <a href="/profile/${username}">
                        <img src="${avatar}" alt="Avatar" class="rounded-circle" style="width:40px;height:40px;object-fit:cover;">
                    </a>
                    <div>
                        <a href="/profile/${username}" style="color:inherit;text-decoration:none;">
                            <div><b>${name}</b></div>
                            <div class="text-muted small">@${username}</div>
                        </a>
                    </div>
                </div>
                ${!isSelf ? `
                    <button class="btn btn-sm ${isFollowing ? 'btn-outline-primary following' : 'btn-primary'} follow-list-btn" data-user-id="${user.id}" data-is-following="${isFollowing}">
                        ${isFollowing ? 'Unfollow' : 'Follow'}
                    </button>
                ` : ''}
            </li>
        `;
    }

    function renderFollowersList(users, emptyMsg) {
        if (!Array.isArray(users) || !users.length) {
            $followersModalContent.html(`<div class="text-center my-4" style="color: #E7E9EA;">${emptyMsg}</div>`);
            return;
        }
        const html = users.map(user => userListItem(user, user.is_following)).join('');
        $followersModalContent.html(`<ul class="list-group list-group-flush">${html}</ul>`);
        users.forEach(user => {
            $.get(`/users/isFollowing/${user.id}`, function(data) {
                const $btn = $followersModalContent.find(`button[data-user-id="${user.id}"]`);
                updateFollowListBtn($btn, data.is_following);
            });
        });
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
