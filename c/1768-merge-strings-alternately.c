#include <stdlib.h>

char* mergeAlternately(const char* word1, const char* word2) {
    int len1 = strlen(word1), len2 = strlen(word2);
    int minLen = len1 < len2 ? len1 : len2;
    int totalLen = len1 + len2;
    char* result = (char*)malloc(totalLen + 1);
    if (!result) return NULL; 

    int i, j = 0;
    for (i = 0; i < minLen; i++) {
        result[j++] = word1[i];
        result[j++] = word2[i];
    }
    while (i < len1) result[j++] = word1[i++];
    while (i < len2) result[j++] = word2[i++];

    result[j] = '\0';
    return result;
}