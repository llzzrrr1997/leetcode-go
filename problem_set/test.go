package problemset

import "math"

func minCommitter(files [][]int) int {
	fileList := make([]int, len(files))
	reviewerFilesMap := make([][]int, 10)
	reviewed := make([]bool, 10)
	for i := range files {
		for _, reviewer := range files[i] {
			reviewerFilesMap[reviewer] = append(reviewerFilesMap[reviewer], i)
		}
	}
	min := math.MaxInt
	backtrack(reviewerFilesMap, fileList, 0, nil, reviewed, &min)
	return min
}

func backtrack(reviewerFilesMap [][]int, fileList []int, fileCnt int, curReviewer []int, reviewed []bool, min *int) {
	if fileCnt == len(fileList) {
		if *min > len(curReviewer) {
			*min = len(curReviewer)
		}
		return
	}
	for i := range reviewerFilesMap {
		if reviewed[i] {
			continue
		}
		reviewed[i] = true
		curReviewer = append(curReviewer, i)
		for _, fileIndex := range reviewerFilesMap[i] {
			fileList[fileIndex]++
			if fileList[fileIndex] == 1 {
				fileCnt++
			}
		}
		backtrack(reviewerFilesMap, fileList, fileCnt, curReviewer, reviewed, min)
		reviewed[i] = false
		curReviewer = curReviewer[:len(curReviewer)-1]
		for _, fileIndex := range reviewerFilesMap[i] {
			fileList[fileIndex]--
			if fileList[fileIndex] == 0 {
				fileCnt--
			}
		}
	}
}
