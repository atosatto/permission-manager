package resources

import (
	"context"

	k8sclient "k8s.io/client-go/kubernetes"
)

// ResourceService allows to list and manage the life-cycle
// of the various K8s cluster resources managed by the PermissionManager.
type ResourcesService interface {
	UserService
	GetAllNamespaces(ctx context.Context) (names []string, err error)
}

// resourceService implements the ResourceService interface.
type resourcesService struct {
	kubeclient k8sclient.Interface
}

// NewResourcesService returns a new instance of a ResourceService
// allowing to interact with a K8s cluster via the given K8s client interface.
func NewResourcesService(kc k8sclient.Interface) ResourcesService {
	return &resourcesService{
		kubeclient: kc,
	}
}
