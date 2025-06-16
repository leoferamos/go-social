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
                const postHtml = `<article class="feed-post">
                    <div class="profile-pic-placeholder"></div>
                    <div class="post-content-area">
                        <header class="post-header">
                            <span class="post-author-name">${escapeHtml(post.author_username)}</span>
                            <span class="post-author-username">@${escapeHtml(post.author_username)}</span>
                            <span class="post-time" data-time="${escapeHtml(post.created_at)}" title=""></span>
                        </header>
                        <p class="post-text">${escapeHtml(post.content)}</p>
                        <footer class="post-actions">
                            <div><i class="bi bi-heart"></i> <span>${post.likes}</span></div>
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

    function escapeHtml(text) {
        return text
            .replace(/&/g, "&amp;")
            .replace(/</g, "&lt;")
            .replace(/>/g, "&gt;")
            .replace(/"/g, "&quot;")
            .replace(/'/g, "&#039;");
    }
});
