package xilution

import (
	b64 "encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dgrijalva/jwt-go"
)

func buildTestId() string {
	return strings.Join(strings.Split(gofakeit.UUID(), "-"), "")
}

func buildJwtToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})
	tokenString, _ := token.SignedString([]byte("AllYourBase"))

	return tokenString
}

func buildTestOrganization() Organization {
	return Organization{
		Type:           "organization",
		ID:             buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		Name:           gofakeit.Company(),
		Logo:           b64.StdEncoding.EncodeToString([]byte(string(gofakeit.ImagePng(50, 50)))),
		Domain:         gofakeit.DomainName(),
		AuthClientId:   buildTestId(),
		OrganizationId: buildTestId(),
		Active:         gofakeit.Bool(),
	}
}

func buildTestClient() Client {
	return Client{
		Type:           "client",
		ID:             buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		Name:           gofakeit.Word(),
		Grants:         []string{"password", "client_credentials", "authorization_code", "refresh_token"},
		RedirectUris:   []string{gofakeit.URL(), gofakeit.URL(), gofakeit.URL()},
		OrganizationId: buildTestId(),
		Active:         gofakeit.Bool(),
	}
}

func buildTestUser() User {
	return User{
		Type:           "user",
		ID:             buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		FirstName:      gofakeit.FirstName(),
		LastName:       gofakeit.LastName(),
		Email:          gofakeit.Email(),
		Username:       gofakeit.Username(),
		OrganizationId: buildTestId(),
		Active:         gofakeit.Bool(),
	}
}

func buildTestGitAccount() GitAccount {
	return GitAccount{
		Type:           "git-account",
		ID:             buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		Provider:       "GIT_HUB",
		Name:           gofakeit.Word(),
		OrganizationId: buildTestId(),
	}
}

func buildTestGitRepo() GitRepo {
	return GitRepo{
		Type:           "git-repo",
		ID:             buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		Name:           gofakeit.Word(),
		GitAccountId:   buildTestId(),
		OrganizationId: buildTestId(),
	}
}

func buildTestGitRepoEvent() GitRepoEvent {
	return GitRepoEvent{
		Type:         "git-repo-event",
		ID:           buildTestId(),
		OwningUserId: buildTestId(),
		CreatedAt:    gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:   gofakeit.Date().Format(time.RFC3339),
		EventType:    "CREATE_REPO_FROM_TEMPLATE_REPO",
		Parameters: map[string]interface{}{
			"sourceOwner":   gofakeit.Word(),
			"sourceRepo":    gofakeit.Word(),
			"rootPath":      "/",
			"targetOwner":   gofakeit.Word(),
			"targetRepo":    gofakeit.Word(),
			"isPrivate":     true,
			"description":   gofakeit.Sentence(10),
			"commitMessage": gofakeit.Sentence(10),
			"params":        `{"foo": "bar", "boo": "baz"}`,
		},
		GitAccountId:   buildTestId(),
		GitRepoId:      buildTestId(),
		OrganizationId: buildTestId(),
	}
}

func buildTestCloudProvider() CloudProvider {
	return CloudProvider{
		Type:           "cloud-provider",
		ID:             buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		Provider:       "AWS",
		Name:           gofakeit.Word(),
		AccountId:      strconv.Itoa(gofakeit.Number(100000000, 999999999)),
		OrganizationId: buildTestId(),
	}
}

func buildTestVpcPipeline() VpcPipeline {
	return VpcPipeline{
		Type:            "pipeline",
		ID:              buildTestId(),
		CloudProviderId: buildTestId(),
		OwningUserId:    buildTestId(),
		CreatedAt:       gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:      gofakeit.Date().Format(time.RFC3339),
		OrganizationId:  buildTestId(),
	}
}

func buildTestK8sPipeline() K8sPipeline {
	return K8sPipeline{
		Type:           "pipeline",
		ID:             buildTestId(),
		VpcPipelineId:  buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		OrganizationId: buildTestId(),
	}
}

func buildTestWordPressPipeline() WordPressPipeline {
	return WordPressPipeline{
		Type:          "pipeline",
		ID:            buildTestId(),
		K8sPipelineId: buildTestId(),
		GitRepoId:     buildTestId(),
		Branch:        gofakeit.Word(),
		Stages: []Stage{
			{
				Name: gofakeit.Word(),
			},
			{
				Name: gofakeit.Word(),
			},
			{
				Name: gofakeit.Word(),
			},
		},
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		OrganizationId: buildTestId(),
	}
}

func buildTestStaticContentPipeline() StaticContentPipeline {
	return StaticContentPipeline{
		Type:            "pipeline",
		ID:              buildTestId(),
		CloudProviderId: buildTestId(),
		GitRepoId:       buildTestId(),
		Branch:          gofakeit.Word(),
		Stages: []Stage{
			{
				Name: gofakeit.Word(),
			},
			{
				Name: gofakeit.Word(),
			},
			{
				Name: gofakeit.Word(),
			},
		},
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		OrganizationId: buildTestId(),
	}
}

func buildTestApiPipeline() ApiPipeline {
	return ApiPipeline{
		Type:          "pipeline",
		ID:            buildTestId(),
		VpcPipelineId: buildTestId(),
		GitRepoId:     buildTestId(),
		Branch:        gofakeit.Word(),
		Stages: []Stage{
			{
				Name: gofakeit.Word(),
			},
			{
				Name: gofakeit.Word(),
			},
			{
				Name: gofakeit.Word(),
			},
		},
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		OrganizationId: buildTestId(),
	}
}
