package mailer

import "testing"

type MockMailer struct {
	To              string
	Subject         string
	Body            string
	AttachmentPaths []string
	SendErr         error
	Called          bool
}

func (m *MockMailer) SendEmail(to string, subject string, body string, attachmentPaths ...string) error {
	m.To = to
	m.Subject = subject
	m.Body = body
	m.AttachmentPaths = attachmentPaths
	m.Called = true
	return m.SendErr
}

func TestMockMailer(t *testing.T) {
	mock := &MockMailer{}
	err := mock.SendEmail("test@example.com", "Test Subject", "Test Body", "cv.pdf", "portfolio.pptx")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !mock.Called {
		t.Error("expected MockMailer to be marked called")
	}
	if mock.To != "test@example.com" {
		t.Errorf("expected To test@example.com, got %s", mock.To)
	}
	if mock.Subject != "Test Subject" {
		t.Errorf("expected Subject 'Test Subject', got '%s'", mock.Subject)
	}
	if mock.Body != "Test Body" {
		t.Errorf("expected Body 'Test Body', got '%s'", mock.Body)
	}
	if len(mock.AttachmentPaths) != 2 || mock.AttachmentPaths[0] != "cv.pdf" || mock.AttachmentPaths[1] != "portfolio.pptx" {
		t.Errorf("expected AttachmentPaths [cv.pdf portfolio.pptx], got %v", mock.AttachmentPaths)
	}
}
