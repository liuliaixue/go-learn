`git config --global http.proxy http://127.0.0.1:1080
git config --global https.proxy https://127.0.0.1:1080
`
`
git config --global --unset http.proxy
git config --global --unset https.proxy
`

git config --global http.https://github.com.proxy socks5://127.0.0.1:1080
git config --global https.https://github.com.proxy socks5://127.0.0.1:1080

git config --global --unset http.https://github.com.proxy
git config --global --unset https.https://github.com.proxy
