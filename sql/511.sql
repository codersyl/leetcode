# MySQL
# leetcode 511
# https://leetcode-cn.com/problems/game-play-analysis-i/

SELECT
    player_id,
    MIN(event_date) AS first_login
FROM Activity
GROUP BY player_id;