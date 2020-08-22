class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        result=[]
        for index,letter in enumerate(strs[0]):
            print(strs[1:])
            for word in strs[1:]:
                print(word)
                print(index-1)
                print(len(word))
                if index-1<=len(word) or letter != word[index]:
                    break
                break
            result.append(letter)
        print(result)

if __name__ == "__main__":
    a=Solution()
    test_case=["flower","flow","flight"]
    a.longestCommonPrefix(test_case)