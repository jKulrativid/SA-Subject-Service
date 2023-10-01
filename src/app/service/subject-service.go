package service

import (
	"context"
	"fmt"

	"github.com/jKulrativid/SA-Subject-Service/src/app/entity"
	"github.com/jKulrativid/SA-Subject-Service/src/app/repository"
	pb "github.com/jKulrativid/SA-Subject-Service/src/grpc/subject"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SubjectService struct {
	pb.UnimplementedSubjectServiceServer
	subjectRepo repository.SubjectRepository
}

func NewSubjectService(subjectRepo repository.SubjectRepository) *SubjectService {
	return &SubjectService{subjectRepo: subjectRepo}
}

func SubjectToPb(subject *entity.Subject) *pb.Subject {
	sectionIds := make([]int64, 0)
	for _, section := range subject.Sections {
		sectionIds = append(sectionIds, section.Id)
	}

	return &pb.Subject{
		Id:          subject.Id,
		SubjectId:   subject.SubjectId,
		Name:        subject.Name,
		Semester:    subject.Semester,
		SectionIds:  sectionIds,
		Year:        subject.Year,
		Faculty:     subject.Faculty,
		Description: subject.Description,
	}
}

func (s *SubjectService) PaginateSubjects(ctx context.Context, req *pb.PaginateSubjectRequest) (*pb.PaginateSubjectResponse, error) {
	if req.PageNumber < 1 {
		return nil, status.Error(codes.InvalidArgument, "page number must be a positive integer")
	}

	query := make(map[string]interface{})

	if req.SubjectId != "" {
		query["subject_id"] = req.SubjectId
	}
	if req.Name != "" {
		query["name"] = req.Name
	}
	if len(req.SemesterWhitelist) != 0 {
		query["semester_whitelist"] = req.SemesterWhitelist
	}
	if req.YearRangeStart != 0 && req.YearRangeStop != 0 {
		isValidYearRange := req.YearRangeStart <= req.YearRangeStop
		if isValidYearRange {
			query["year_range_start"] = req.YearRangeStart
			query["year_range_stop"] = req.YearRangeStop
		}
	}

	metadata, subjects, err := s.subjectRepo.PaginateSubjects(req.PageNumber, query)
	if err != nil {
		return nil, err
	}

	resp := pb.PaginateSubjectResponse{
		PageNumber: metadata.Page,
		PerPage:    metadata.PerPage,
		PageCount:  metadata.PageCount,
		TotalCount: metadata.TotalCount,
		Subjects:   make([]*pb.SubjectMetadata, 0),
	}

	for _, subject := range subjects {
		resp.Subjects = append(resp.Subjects, &pb.SubjectMetadata{
			Id:        subject.Id,
			Name:      subject.Name,
			SubjectId: subject.SubjectId,
			Semester:  subject.Semester,
			Year:      subject.Year,
		})
	}

	return &resp, nil
}

func (s *SubjectService) GetSubjectById(ctx context.Context, req *pb.GetSubjectByIdRequest) (*pb.GetSubjectByIdResponse, error) {
	subject, err := s.subjectRepo.FindSubjectById(req.Id)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return nil, status.Error(codes.NotFound, "not found")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.GetSubjectByIdResponse{Subject: SubjectToPb(subject)}, nil
}

func (s *SubjectService) CreateSubject(ctx context.Context, req *pb.CreateSubjectRequest) (*pb.CreateSubjectResponse, error) {
	if req.SubjectId == "" {
		return nil, status.Error(codes.InvalidArgument, "subject ID not provided")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "subject name not provided")
	}
	if req.Semester < 1 || req.Semester > 3 {
		return nil, status.Error(codes.InvalidArgument, "subject semester should be 1,2 or 3")
	}
	if req.Year < 2000 {
		return nil, status.Error(codes.InvalidArgument, "subject year must be after 1999")
	}
	if req.Faculty == "" {
		return nil, status.Error(codes.InvalidArgument, "subject faculty not provided")
	}

	subject := entity.Subject{
		SubjectId:   req.SubjectId,
		Name:        req.Name,
		Semester:    req.Semester,
		Year:        req.Year,
		Faculty:     req.Faculty,
		Description: req.Description,
	}

	if err := s.subjectRepo.CreateSubject(&subject); err != nil {
		fmt.Println(err)
		switch err {
		case entity.ErrConstraintViolation:
			return nil, status.Error(codes.InvalidArgument, "bad request")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.CreateSubjectResponse{Subject: SubjectToPb(&subject)}, nil
}

func (s *SubjectService) UpdateSubject(ctx context.Context, req *pb.UpdateSubjectRequest) (*pb.UpdateSubjectResponse, error) {
	if req.Semester != 0 && (req.Semester < 1 || req.Semester > 3) {
		return nil, status.Error(codes.InvalidArgument, "subject semester should be 1,2 or 3")
	}
	if req.Year != 0 && (req.Year < 2000) {
		return nil, status.Error(codes.InvalidArgument, "subject year must be after 1999")
	}

	subject := entity.Subject{
		Id:          req.Id,
		SubjectId:   req.SubjectId,
		Name:        req.Name,
		Semester:    req.Semester,
		Year:        req.Year,
		Faculty:     req.Faculty,
		Description: req.Description,
	}

	err := s.subjectRepo.UpdateSubject(&subject)
	if err != nil {
		switch err {
		case entity.ErrConstraintViolation:
			return nil, status.Error(codes.InvalidArgument, "bad request")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.UpdateSubjectResponse{Subject: SubjectToPb(&subject)}, nil
}

func (s *SubjectService) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectRequest) (*pb.DeleteSubjectResponse, error) {
	if req.Id < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid ID")
	}

	subject, err := s.subjectRepo.DeleteSubjectById(req.Id)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			return nil, status.Error(codes.NotFound, "subject with given ID not found")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteSubjectResponse{Subject: SubjectToPb(subject)}, nil
}
