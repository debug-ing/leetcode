function mergeAlternately(word1, word2) {
    let result = '';
    let minLen = Math.min(word1.length, word2.length);

    for (let i = 0; i < minLen; i++) {
        result += word1[i] + word2[i];
    }

    result += word1.slice(minLen) + word2.slice(minLen);
    return result;
}