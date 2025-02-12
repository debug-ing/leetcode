merge_alternately() {
    local word1="$1"
    local word2="$2"
    local result=""
    local len1=${#word1}
    local len2=${#word2}
    local min_len=$((len1 < len2 ? len1 : len2))
    for ((i = 0; i < min_len; i++)); do
        result+="${word1:i:1}${word2:i:1}"
    done
    result+="${word1:min_len}"
    result+="${word2:min_len}"
    echo "$result"
}

merge_alternately $1 $2