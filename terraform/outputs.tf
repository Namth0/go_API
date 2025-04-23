output "cluster_endpoint" {
  description = "Endpoit pour le panneau de controle EKS"
  value       = module.eks.cluster_endpoint
}

output "cluster_security_group_id" {
  description = "Security group ids attached to the cluster control plane"
  value       = module.eks.cluster_security_group_id
}

output "region" {
  description = "Region AWS"
  value       = var.region
}

output "cluster_name" {
  description = "Nom du cluster kubernetes"
  value       = module.eks.cluster_name
}
