package internal

import (
	. "../../pull_request"
	. "../../utils"
)

type prRequest struct {
	branch, environment, version string
	filesChanged                 []string
}

func OpenNumberPR(branch, number string, files []string, environment ReleaseEnvironment) error {
	request := prRequest{
		branch:       branch,
		environment:  environment.String(),
		version:      GetBranchWithDotAndNumberWithDashFormat(branch, number),
		filesChanged: files,
	}
	return openNumberPR(&request)
}

func openNumberPR(prReq *prRequest) error {
	pr, err := NewPullRequestForNumber(prReq)
	if err != nil {
		return err
	}
	return pr.Open()
}

func (prReq *prRequest) Branch() string {
	return prReq.branch
}

func (prReq *prRequest) Environment() string {
	return prReq.environment
}

func (prReq *prRequest) Version() string {
	return prReq.version
}

func (prReq *prRequest) FilePaths() []string {
	return prReq.filesChanged
}
