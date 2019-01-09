package main

import (
	"encoding/json"
	"strconv"

	"github.com/HcashOrg/hcAutonomy/HcashOrgplugin"
	pd "github.com/HcashOrg/hcAutonomy/hcAutonomyd/api/v1"
	www "github.com/HcashOrg/hcAutonomy/hcAutonomywww/api/v1"
)

func convertCastVoteReplyFromHcPlugin(cvr HcashOrgplugin.CastVoteReply) www.CastVoteReply {
	return www.CastVoteReply{
		ClientSignature: cvr.ClientSignature,
		Signature:       cvr.Signature,
		Error:           cvr.Error,
	}
}

func convertCastVoteFromWWW(b www.CastVote) HcashOrgplugin.CastVote {
	return HcashOrgplugin.CastVote{
		Token:     b.Token,
		Ticket:    b.Ticket,
		VoteBit:   b.VoteBit,
		Signature: b.Signature,
	}
}

func convertBallotFromWWW(b www.Ballot) HcashOrgplugin.Ballot {
	br := HcashOrgplugin.Ballot{
		Votes: make([]HcashOrgplugin.CastVote, 0, len(b.Votes)),
	}
	for _, v := range b.Votes {
		br.Votes = append(br.Votes, convertCastVoteFromWWW(v))
	}
	return br
}

func convertBallotReplyFromHcPlugin(b HcashOrgplugin.BallotReply) www.BallotReply {
	br := www.BallotReply{
		Receipts: make([]www.CastVoteReply, 0, len(b.Receipts)),
	}
	for _, v := range b.Receipts {
		br.Receipts = append(br.Receipts,
			convertCastVoteReplyFromHcPlugin(v))
	}
	return br
}

func convertVoteOptionFromWWW(vo www.VoteOption) HcashOrgplugin.VoteOption {
	return HcashOrgplugin.VoteOption{
		Id:          vo.Id,
		Description: vo.Description,
		Bits:        vo.Bits,
	}
}

func convertVoteOptionsFromWWW(vo []www.VoteOption) []HcashOrgplugin.VoteOption {
	vor := make([]HcashOrgplugin.VoteOption, 0, len(vo))
	for _, v := range vo {
		vor = append(vor, convertVoteOptionFromWWW(v))
	}
	return vor
}

func convertVoteFromWWW(v www.Vote) HcashOrgplugin.Vote {
	return HcashOrgplugin.Vote{
		Token:            v.Token,
		Mask:             v.Mask,
		Duration:         v.Duration,
		QuorumPercentage: v.QuorumPercentage,
		PassPercentage:   v.PassPercentage,
		Options:          convertVoteOptionsFromWWW(v.Options),
	}
}

func convertAuthorizeVoteFromWWW(av www.AuthorizeVote) HcashOrgplugin.AuthorizeVote {
	return HcashOrgplugin.AuthorizeVote{
		Action:    av.Action,
		Token:     av.Token,
		PublicKey: av.PublicKey,
		Signature: av.Signature,
	}
}

func convertAuthorizeVoteReplyFromHcplugin(avr HcashOrgplugin.AuthorizeVoteReply) www.AuthorizeVoteReply {
	return www.AuthorizeVoteReply{
		Action:  avr.Action,
		Receipt: avr.Receipt,
	}
}

func convertStartVoteFromWWW(sv www.StartVote) HcashOrgplugin.StartVote {
	return HcashOrgplugin.StartVote{
		PublicKey: sv.PublicKey,
		Vote:      convertVoteFromWWW(sv.Vote),
		Signature: sv.Signature,
	}
}

func convertStartVoteFromHcplugin(sv HcashOrgplugin.StartVote) www.StartVote {
	return www.StartVote{
		PublicKey: sv.PublicKey,
		Vote:      convertVoteFromHcplugin(sv.Vote),
		Signature: sv.Signature,
	}
}

func convertStartVoteReplyFromHcplugin(svr HcashOrgplugin.StartVoteReply) www.StartVoteReply {
	return www.StartVoteReply{
		StartBlockHeight: svr.StartBlockHeight,
		StartBlockHash:   svr.StartBlockHash,
		EndHeight:        svr.EndHeight,
		EligibleTickets:  svr.EligibleTickets,
	}
}

func convertVoteOptionFromHcplugin(vo HcashOrgplugin.VoteOption) www.VoteOption {
	return www.VoteOption{
		Id:          vo.Id,
		Description: vo.Description,
		Bits:        vo.Bits,
	}
}

func convertVoteOptionsFromHcplugin(vo []HcashOrgplugin.VoteOption) []www.VoteOption {
	vor := make([]www.VoteOption, 0, len(vo))
	for _, v := range vo {
		vor = append(vor, convertVoteOptionFromHcplugin(v))
	}
	return vor
}

func convertVoteFromHcplugin(v HcashOrgplugin.Vote) www.Vote {
	return www.Vote{
		Token:            v.Token,
		Mask:             v.Mask,
		Duration:         v.Duration,
		QuorumPercentage: v.QuorumPercentage,
		PassPercentage:   v.PassPercentage,
		Options:          convertVoteOptionsFromHcplugin(v.Options),
	}
}

func convertCastVoteFromHcplugin(cv HcashOrgplugin.CastVote) www.CastVote {
	return www.CastVote{
		Token:     cv.Token,
		Ticket:    cv.Ticket,
		VoteBit:   cv.VoteBit,
		Signature: cv.Signature,
	}
}

func convertCastVotesFromHcplugin(cv []HcashOrgplugin.CastVote) []www.CastVote {
	cvr := make([]www.CastVote, 0, len(cv))
	for _, v := range cv {
		cvr = append(cvr, convertCastVoteFromHcplugin(v))
	}
	return cvr
}

func convertVoteResultsReplyFromHcplugin(vrr HcashOrgplugin.VoteResultsReply, ir inventoryRecord) www.VoteResultsReply {
	return www.VoteResultsReply{
		StartVote:      convertStartVoteFromHcplugin(vrr.StartVote),
		CastVotes:      convertCastVotesFromHcplugin(vrr.CastVotes),
		StartVoteReply: ir.voting,
	}
}

func convertPropStatusFromWWW(s www.PropStatusT) pd.RecordStatusT {
	switch s {
	case www.PropStatusNotFound:
		return pd.RecordStatusNotFound
	case www.PropStatusNotReviewed:
		return pd.RecordStatusNotReviewed
	case www.PropStatusCensored:
		return pd.RecordStatusCensored
	case www.PropStatusPublic:
		return pd.RecordStatusPublic
	case www.PropStatusAbandoned:
		return pd.RecordStatusArchived
	}
	return pd.RecordStatusInvalid
}

func convertPropFileFromWWW(f www.File) pd.File {
	return pd.File{
		Name:    f.Name,
		MIME:    f.MIME,
		Digest:  f.Digest,
		Payload: f.Payload,
	}
}

func convertPropFilesFromWWW(f []www.File) []pd.File {
	files := make([]pd.File, 0, len(f))
	for _, v := range f {
		files = append(files, convertPropFileFromWWW(v))
	}
	return files
}

func convertPropCensorFromWWW(f www.CensorshipRecord) pd.CensorshipRecord {
	return pd.CensorshipRecord{
		Token:     f.Token,
		Merkle:    f.Merkle,
		Signature: f.Signature,
	}
}

// convertPropFromWWW converts a www proposal to a hcAutonomyd record.  This
// function should only be used in tests. Note that convertPropFromWWW can not
// emulate MD properly.
func convertPropFromWWW(p www.ProposalRecord) pd.Record {
	return pd.Record{
		Status:    convertPropStatusFromWWW(p.Status),
		Timestamp: p.Timestamp,
		Metadata: []pd.MetadataStream{{
			ID:      pd.MetadataStreamsMax + 1, // fail deliberately
			Payload: "invalid payload",
		}},
		Files:            convertPropFilesFromWWW(p.Files),
		CensorshipRecord: convertPropCensorFromWWW(p.CensorshipRecord),
	}
}

func convertPropsFromWWW(p []www.ProposalRecord) []pd.Record {
	pr := make([]pd.Record, 0, len(p))
	for _, v := range p {
		pr = append(pr, convertPropFromWWW(v))
	}
	return pr
}

///////////////////////////////
func convertPropStatusFromPD(s pd.RecordStatusT) www.PropStatusT {
	switch s {
	case pd.RecordStatusNotFound:
		return www.PropStatusNotFound
	case pd.RecordStatusNotReviewed:
		return www.PropStatusNotReviewed
	case pd.RecordStatusCensored:
		return www.PropStatusCensored
	case pd.RecordStatusPublic:
		return www.PropStatusPublic
	case pd.RecordStatusUnreviewedChanges:
		return www.PropStatusUnreviewedChanges
	case pd.RecordStatusArchived:
		return www.PropStatusAbandoned
	}
	return www.PropStatusInvalid
}

func convertPropFileFromPD(f pd.File) www.File {
	return www.File{
		Name:    f.Name,
		MIME:    f.MIME,
		Digest:  f.Digest,
		Payload: f.Payload,
	}
}

func convertPropFilesFromPD(f []pd.File) []www.File {
	files := make([]www.File, 0, len(f))
	for _, v := range f {
		files = append(files, convertPropFileFromPD(v))
	}
	return files
}

func convertPropCensorFromPD(f pd.CensorshipRecord) www.CensorshipRecord {
	return www.CensorshipRecord{
		Token:     f.Token,
		Merkle:    f.Merkle,
		Signature: f.Signature,
	}
}

func convertPropFromPD(p pd.Record) www.ProposalRecord {
	md := &BackendProposalMetadata{}
	var statusChangeMsg string
	for _, v := range p.Metadata {
		if v.ID == mdStreamGeneral {
			m, err := decodeBackendProposalMetadata([]byte(v.Payload))
			if err != nil {
				log.Errorf("could not decode metadata '%v' token '%v': %v",
					p.Metadata, p.CensorshipRecord.Token, err)
				break
			}
			md = m
		}

		if v.ID == mdStreamChanges {
			var mdc MDStreamChanges
			err := json.Unmarshal([]byte(v.Payload), &mdc)
			if err != nil {
				break
			}
			statusChangeMsg = mdc.StatusChangeMessage
		}
	}

	var state www.PropStateT
	status := convertPropStatusFromPD(p.Status)
	switch status {
	case www.PropStatusNotReviewed, www.PropStatusUnreviewedChanges,
		www.PropStatusCensored:
		state = www.PropStateUnvetted
	case www.PropStatusPublic, www.PropStatusAbandoned:
		state = www.PropStateVetted
	default:
		state = www.PropStateInvalid
	}

	return www.ProposalRecord{
		Name:                md.Name,
		State:               state,
		Status:              status,
		Timestamp:           md.Timestamp,
		PublicKey:           md.PublicKey,
		Signature:           md.Signature,
		Files:               convertPropFilesFromPD(p.Files),
		CensorshipRecord:    convertPropCensorFromPD(p.CensorshipRecord),
		Version:             p.Version,
		StatusChangeMessage: statusChangeMsg,
	}
}

func convertErrorStatusFromPD(s int) www.ErrorStatusT {
	switch pd.ErrorStatusT(s) {
	case pd.ErrorStatusInvalidFileDigest:
		return www.ErrorStatusInvalidFileDigest
	case pd.ErrorStatusInvalidBase64:
		return www.ErrorStatusInvalidBase64
	case pd.ErrorStatusInvalidMIMEType:
		return www.ErrorStatusInvalidMIMEType
	case pd.ErrorStatusUnsupportedMIMEType:
		return www.ErrorStatusUnsupportedMIMEType
	case pd.ErrorStatusInvalidRecordStatusTransition:
		return www.ErrorStatusInvalidPropStatusTransition
	case pd.ErrorStatusInvalidFilename:
		return www.ErrorStatusInvalidFilename

		// These cases are intentionally omitted because
		// they are indicative of some internal server error,
		// so ErrorStatusInvalid is returned.
		//
		//case pd.ErrorStatusInvalidRequestPayload
		//case pd.ErrorStatusInvalidChallenge
	}
	return www.ErrorStatusInvalid
}

func convertVoteResultsFromHcplugin(vrr HcashOrgplugin.VoteResultsReply) []www.VoteOptionResult {
	// counter of votes received
	var vr uint64
	var ors []www.VoteOptionResult
	for _, o := range vrr.StartVote.Vote.Options {
		vr = 0
		for _, v := range vrr.CastVotes {
			vb, err := strconv.ParseUint(v.VoteBit, 10, 64)
			if err != nil {
				log.Infof("it shouldn't happen")
				continue
			}
			if vb == o.Bits {
				vr++
			}
		}

		// append to vote options result slice
		ors = append(ors, www.VoteOptionResult{
			VotesReceived: vr,
			Option:        convertVoteOptionFromHcplugin(o),
		})
	}
	return ors
}
