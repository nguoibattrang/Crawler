CUR_DIR=$(pwd)
OUTPUT=$CUR_DIR/output

function craw_jira() {
    echo "Crawling $1"
    curl "https://jira.viettelcyber.com/browse/$1" \
    -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' \
    -H 'Accept-Language: vi-VN,vi;q=0.9,fr-FR;q=0.8,fr;q=0.7,en-US;q=0.6,en;q=0.5' \
    -H 'Cache-Control: max-age=0' \
    -H 'Connection: keep-alive' \
    -H 'Cookie: jira.editor.user.mode=wysiwyg; RETURN_TO_COOKIE=%2Fbrowse%2FVSAFE-1; JSESSIONID=0322362281E2DDEAEE2450A5AC4C7339; mo.jira-oauth.logoutcookie=10837907-2dcf-4d2c-9705-7c58f683b674; mo.jira-oauth.baseurl=https://jira.viettelcyber.com; atlassian.xsrf.token=BACD-QJMD-4NA4-CI1L_d5b53218d5e9207cd8348e8497a77835de24cc7f_lin' \
    -H 'Sec-Fetch-Dest: document' \
    -H 'Sec-Fetch-Mode: navigate' \
    -H 'Sec-Fetch-Site: same-origin' \
    -H 'Sec-Fetch-User: ?1' \
    -H 'Upgrade-Insecure-Requests: 1' \
    -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36' \
    -H 'sec-ch-ua: "Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"' \
    -H 'sec-ch-ua-mobile: ?0' \
    -H 'sec-ch-ua-platform: "Windows"' \
    -o output/$1.html -s
}

function start() {
    mkdir -p $OUTPUT
    i=$1
    while [ $i -lt $2 ]
    do
        craw_file=vsafe-$i
        craw_jira $craw_file
        i=$(( $i + 1 ))
    done;
}

function clean() {
    rm -rf $OUTPUT
}

function usage() {
    echo "Usage:"
    echo "sh jira_crawler.sh [start_topic] [end_topic]"
}

$@

if [ -z $1 ];
then
    usage
fi