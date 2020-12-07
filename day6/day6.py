# Count the distinct letters in a group separated by a blank line
# abcx
# abcy
# abcz
#
# 6
# Print the sum of distinct counts
#
# Part b
# Count letters that appear in every entry in a group, sum the counts

import sys

def countGroup(groupLines):
    answers = set()
    for line in groupLines:
        answerLine = line
        if line[-1] == '\n':
            answerLine = line[:-1]
        for answer in answerLine:
            answers.add(answer)
    return len(answers)

def countGroupPartTwo(groupLines):
    answers = set()
    for (i, line) in enumerate(groupLines):
        answerLine = line
        if line[-1] == '\n':
            answerLine = line[:-1]
        if i == 0:
            answers = set(answerLine)
        else:
            answers &= set(answerLine)
    return len(answers)

if __name__ == "__main__":
    group = []
    countSum = 0
    countSumPtTwo = 0
    for line in sys.stdin: 
        if line == '\n':
            countSum += countGroup(group)
            countSumPtTwo += countGroupPartTwo(group)
            group = []
        else:
            group.append(line)
    countSum += countGroup(group)
    countSumPtTwo += countGroupPartTwo(group)
    print(countSum)
    print(countSumPtTwo)