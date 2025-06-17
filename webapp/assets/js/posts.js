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
            success: function(post) {
                $('.create-post-textarea').val('');
                const postHtml = `<article class="feed-post" data-post-id="${post.id}">
                    <div class="profile-pic-placeholder"></div>
                    <div class="post-content-area">
                        <header class="post-header">
                            <span class="post-author-name">${escapeHtml(post.author_username)}</span>
                            <span class="post-author-username">@${escapeHtml(post.author_username)}</span>
                            <span class="post-time" data-time="${escapeHtml(post.created_at)}" title=""></span>
                        </header>
                        <p class="post-text">${escapeHtml(post.content)}</p>
                        <footer class="post-actions">
                            <div>
                                <i class="bi bi-heart like-post"></i>
                                <span>${post.likes}</span>
                            </div>
                        </footer>
                    </div>
                </article>`;
                $('#feed-posts').prepend(postHtml);
                document.querySelectorAll('.post-time').forEach(function(span) {
                    const iso = span.getAttribute('data-time');
                    if (iso) {
                        span.textContent = formatRelativeDate(iso);
                        span.title = formatFullDate(iso);
                    }
                });
            },
            error: function(xhr) {
                let msg = 'Error creating post. Please try again.';
                if (xhr && xhr.responseText) {
                    try {
                        const resp = JSON.parse(xhr.responseText);
                        if (resp && resp.error) {
                            msg = 'Erro: ' + resp.error;
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

    // Like/Unlike handler
    $('#feed-posts').on('click', '.like-post', function() {
        const $icon = $(this);
        const $post = $icon.closest('.feed-post');
        const postId = $post.data('post-id');
        const liked = $icon.hasClass('bi-heart-fill');

        // Escolhe a rota correta
        const url = liked ? `/posts/${postId}/unlike` : `/posts/${postId}/like`;

        $.ajax({
            url: url,
            method: 'POST',
            success: function(data) {
                // data esperado: { id, likes, liked_by_me }
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

    $('#feed-posts').on('click', '.delete-post-btn', function(e) {
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
