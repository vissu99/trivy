package all

import (
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/config/azurearm"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/config/cloudformation"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/config/dockerfile"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/config/helm"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/config/k8s"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/config/terraform"
	_ "github.com/aquasecurity/trivy/pkg/fanal/analyzer/config/terraformplan"
)
