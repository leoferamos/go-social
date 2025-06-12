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
                            <span class="post-author-name">${post.author_username}</span>
                            <span class="post-author-username">@${post.author_username}</span>
                            <span class="post-time" data-time="${post.created_at}" title=""></span>
                        </header>
                        <p class="post-text">${post.content}</p>
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
            error: function() {
                alert('Error creating post. Please try again.');
            }
        });
    }
});
