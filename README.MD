# How to use

1- download release according to your OS

## init config

```
onydev init --clientID=clientIDToUse
--realm=realmToUse
--keycloakURL=https://keycloak.example.com --onboardingURL=https://onboarding.example.com
```

will generate a config file for onydev in $Home/.onydev/config.yaml

## get kubeconfig

```
onydev get kubeconfig --generate
```

will create a .kube/config in $HOME

then kubectl get pods :)

## get token

Onydev is also able to retrieve token on a specified realm client and keycloak server by default it use the one which was configured on init :

```
onydev auth token
```

To override :

```
onydev auth token --clientID=clientID --keycloakURL=keycloakURL --realm=realm
```
