package server

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// PodSandboxStatus returns the Status of the PodSandbox.
func (s *Server) PodSandboxStatus(ctx context.Context, req *pb.PodSandboxStatusRequest) (resp *pb.PodSandboxStatusResponse, err error) {
	sb, err := s.getPodSandboxFromRequest(req.PodSandboxId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "could not find pod %q: %v", req.PodSandboxId, err)
	}

	rStatus := pb.PodSandboxState_SANDBOX_NOTREADY
	if sb.Ready(true) {
		rStatus = pb.PodSandboxState_SANDBOX_READY
	}

	linux := &pb.LinuxPodSandboxStatus{
		Namespaces: &pb.Namespace{
			Options: &pb.NamespaceOption{
				Network: sb.NamespaceOptions().GetNetwork(),
				Ipc:     sb.NamespaceOptions().GetIpc(),
				Pid:     sb.NamespaceOptions().GetPid(),
			},
		},
	}

	sandboxID := sb.ID()
	resp = &pb.PodSandboxStatusResponse{
		Status: &pb.PodSandboxStatus{
			Id:          sandboxID,
			CreatedAt:   sb.CreatedAt().UnixNano(),
			Network:     &pb.PodSandboxNetworkStatus{},
			State:       rStatus,
			Labels:      sb.Labels(),
			Annotations: sb.Annotations(),
			Metadata:    sb.Metadata(),
			Linux:       linux,
		},
	}

	if len(sb.IPs()) > 0 {
		resp.Status.Network.Ip = sb.IPs()[0]
	}
	if len(sb.IPs()) > 1 {
		resp.Status.Network.AdditionalIps = toPodIPs(sb.IPs()[1:])
	}
	return resp, nil
}

func toPodIPs(ips []string) (result []*pb.PodIP) {
	for _, ip := range ips {
		result = append(result, &pb.PodIP{Ip: ip})
	}
	return result
}
