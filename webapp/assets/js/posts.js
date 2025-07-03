$(document).ready(function() {
    $('.create-post-container').on('submit', createPost);

    function createPost(event) {
        event.preventDefault();
        const content = $('.create-post-textarea').val();
        if (!content.trim()) return;
        $.ajax({
            url: '/posts',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({ content }),
            success: function(html) {
                $('.create-post-textarea').val('');
                $('#feed-posts').prepend(html);
                document.querySelectorAll('.post-time').forEach(function(span) {
                    const iso = span.getAttribute('data-time');
                    if (iso) {
                        span.textContent = formatRelativeDate(iso);
                        span.title = formatFullDate(iso);
                    }
                });
                $('#no-posts-message').remove();
            },
            error: function(xhr) {
                let msg = 'Error creating post. Please try again.';
                if (xhr && xhr.responseText) {
                    try {
                        const resp = JSON.parse(xhr.responseText);
                        if (resp && resp.error) {
                            msg = 'Error: ' + resp.error;
                        }
                    } catch (e) {
                        msg = 'HTTP ' + xhr.status + ': ' + xhr.statusText + '\n' + xhr.responseText;
                    }
                } else if (xhr && xhr.status) {
                    msg = 'HTTP ' + xhr.status + ': ' + xhr.statusText;
                }
                alert(msg);
            }
        });
    }


    $('#profile-posts, #feed-posts').on('click', '.like-post', function() {
        const $icon = $(this);
        const $post = $icon.closest('.feed-post');
        const postId = $post.data('post-id');
        const liked = $icon.hasClass('bi-heart-fill');


        const url = liked ? `/posts/${postId}/unlike` : `/posts/${postId}/like`;

        $.ajax({
            url: url,
            method: 'POST',
            success: function(data) {
                if (data.liked_by_me) {
                    $icon.removeClass('bi-heart').addClass('bi-heart-fill text-danger');
                } else {
                    $icon.removeClass('bi-heart-fill text-danger').addClass('bi-heart');
                }
                $icon.next('span').text(data.likes);
            },
            error: function() {
                alert('Error updating like status.');
            }
        });
    });
    
    let postIdToDelete = null;

    $('#profile-posts, #feed-posts').on('click', '.delete-post-btn', function(e) {
        e.preventDefault();
        postIdToDelete = $(this).data('post-id');
        const deleteModal = new bootstrap.Modal(document.getElementById('deletePostModal'));
        deleteModal.show();
    });

    $('#confirmDeletePostBtn').on('click', function() {
        if (!postIdToDelete) return;
        const $post = $(`.feed-post[data-post-id="${postIdToDelete}"]`);
        $.ajax({
            url: `/posts/${postIdToDelete}`,
            method: 'DELETE',
            success: function() {
                $post.remove();
                postIdToDelete = null;
                const deleteModal = bootstrap.Modal.getInstance(document.getElementById('deletePostModal'));
                deleteModal.hide();
                if ($('#feed-posts .feed-post').length === 0) {
                    $('#feed-posts').append($('#no-posts-message-template').html());
                }
                if ($('#profile-posts .feed-post').length === 0) {
                    $('#profile-posts').append($('#no-posts-message-template').html());
                }
            },
            error: function(xhr) {
                let msg = 'Error deleting post.';
                if (xhr && xhr.responseText) {
                    try {
                        const resp = JSON.parse(xhr.responseText);
                        if (resp && resp.error) {
                            msg = 'Error: ' + resp.error;
                        }
                    } catch (e) {
                        msg = 'HTTP ' + xhr.status + ': ' + xhr.statusText + '\n' + xhr.responseText;
                    }
                } else if (xhr && xhr.status) {
                    msg = 'HTTP ' + xhr.status + ': ' + xhr.statusText;
                }
                alert(msg);
            }
        });
    });

    function escapeHtml(text) {
        return text
            .replace(/&/g, "&amp;")
            .replace(/</g, "&lt;")
            .replace(/>/g, "&gt;")
            .replace(/"/g, "&quot;")
            .replace(/'/g, "&#039;");
    }
});
