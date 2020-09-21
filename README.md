# tfarbe
Inspired from: https://github.com/coinbase/terraform-landscape

Add color to Terraform 12 plan output.

**NOTE**: if you're using terraform <=11, this tool will not help and you should use terraform-landscape.

<img src="./example/screenshot.png" width="65%" alt="Improved Terraform plan output" />

Also formats the ouput for markdown diff.

For example, this:
```
 # module.apps.kubernetes_service.service will be updated in-place
  ~ resource "kubernetes_service" "service" {
        id                    = "some-app"
        load_balancer_ingress = []

      ~ metadata {
            annotations      = {}
            generation       = 0
          ~ labels           = {
                "app.kubernetes.io/managed-by" = "terraform"
              ~ "app.kubernetes.io/version"    = "latest" -> "1.0.1"
            }

  # module.apps.kubernetes_deployment.deployment is tainted, so must be replaced
-/+ resource "kubernetes_deployment" "deployment" {
      ~ id               = "some-app" -> (known after apply)
        wait_for_rollout = true

              ~ spec {
                  - active_deadline_seconds          = 0 -> null
                  - automount_service_account_token  = false -> null
                    dns_policy                       = "ClusterFirst"
                    host_ipc                         = false
                    host_network                     = false
                    host_pid                         = false
                  + hostname                         = (known after apply)
                  + node_name                        = (known after apply)
                    restart_policy                   = "Always"
                  + service_account_name             = (known after apply)
```

becomes:

```diff
# module.apps.kubernetes_service.service will be updated in-place
~   resource "kubernetes_service" "service" {
        id                    = "some-app"
        load_balancer_ingress = []

~       metadata {
            annotations      = {}
            generation       = 0
~           labels           = {
                "app.kubernetes.io/managed-by" = "terraform"
~               "app.kubernetes.io/version"    = "latest" -> "1.0.1"
            }

# module.apps.kubernetes_deployment.deployment is tainted, so must be replaced
-/+ resource "kubernetes_deployment" "deployment" {
~       id               = "some-app" -> (known after apply)
        wait_for_rollout = true

~               spec {
-                   active_deadline_seconds          = 0 -> null
-                   automount_service_account_token  = false -> null
                    dns_policy                       = "ClusterFirst"
                    host_ipc                         = false
                    host_network                     = false
                    host_pid                         = false
+                   hostname                         = (known after apply)
+                   node_name                        = (known after apply)
                    restart_policy                   = "Always"
+                   service_account_name             = (known after apply)
```

## Install

Binaries (should) be available on the [Release page](/releases) for both Linux and Mac. You can simple copy one of these binaries to your PATH.

At the moment it's only been tested in Ubuntu with a small selection of output.

## Usage

```
terraform plan | tfarbe
```

### with Docker

```
docker build . -t tfarbe
terraform plan ... | docker run -i --rm tfarbe
```

## License

This project is released under the [MIT license](LICENSE).
