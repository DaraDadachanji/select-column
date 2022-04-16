# Select-column
`select-column` is a simple command line utility to extract a specific column from a text table passed via stdin

## Why not just use awk, cut, or sed?

Takes longer to type and I don't always remember the full syntax.
I do this enough that it was worth it for me to abstract this specific use-case into it's own command


## Installation

Install Go from the [official website](https://go.dev/)

Clone this repository and build the executable

```Bash
git clone https://github.com/DaraDadachanji/select-column.git
cd select-column
go build
```

Them move it somewhere in your PATH. For example on mac I use:

```Bash
mv ./select-column /usr/local/bin/select-column
```

try it out

```Bash
echo "hello world" | select-column 1
hello
```

## Example usage

I check the pods in a Kubernetes cluster which I'm using to run an Airbyte deployment

```Bash
kubectl get pods
NAME                                   READY   STATUS      RESTARTS   AGE
airbyte-bootloader                     0/1     Completed   0          3d6h
airbyte-db-5b765d55cc-m7749            1/1     Running     0          3d6h
airbyte-minio-5f776f998c-gg6zr         1/1     Running     0          3d6h
airbyte-pod-sweeper-78bfcbd996-j7xsb   1/1     Running     0          3d6h
airbyte-scheduler-7bcb4dbf6-p44tb      1/1     Running     0          3d6h
airbyte-server-7dcd889ff5-h6777        1/1     Running     0          3d6h
airbyte-temporal-5fc5c6f75c-qjqkg      1/1     Running     0          3d6h
airbyte-webapp-6c5f699c69-cb6sk        1/1     Running     0          3d6h
airbyte-worker-7cc86d88db-td79h        1/1     Running     0          3d6h
destination-s3-sync-4-0-oiywd          0/5     Pending     0          10h
destination-s3-sync-4-1-oeyft          0/5     Pending     0          10h
destination-s3-sync-4-2-nnubx          0/5     Pending     0          10h
```

I notice that some of them are pending and have been for some time.
These are hanging pods which will never be launched
because the cluster doesn't have large enough instances to meet the size requirements requested.

I'd like to I'd like to clean up these pods so I grep for 'Pending'

```Bash
kubectl get pods | grep Pending
destination-s3-sync-4-0-oiywd   0/5     Pending   0          10h
destination-s3-sync-4-1-oeyft   0/5     Pending   0          10h
destination-s3-sync-4-2-nnubx   0/5     Pending   0          10h
```

I can now use select-column to cleanly pick just the pod names

```Bash
kubectl get pods | grep Pending | select-column 1
destination-s3-sync-4-0-oiywd
destination-s3-sync-4-1-oeyft
destination-s3-sync-4-2-nnubx
```

and then pipe that to a delete command

```Bash
kubectl get pods | grep Pending | select-column 1 | xargs kubectl delete pod
pod "destination-s3-sync-4-0-oiywd" deleted
pod "destination-s3-sync-4-1-oeyft" deleted
pod "destination-s3-sync-4-2-nnubx" deleted
```

## In reverse

You can pass a negative number to count backwards from the last column

example

```Bash
aws s3 ls | grep salesforce
2022-03-24 16:31:32 irt-dl-us-salesforce-restricted
2022-03-24 14:29:30 irt-dl-us-salesforce-restricted-dev
2022-03-24 16:31:34 irt-dl-us-salesforce-unrestricted
2022-03-24 14:29:30 irt-dl-us-salesforce-unrestricted-dev
```

```Bash
aws s3 ls | grep salesforce | select-column -1
irt-dl-us-salesforce-restricted
irt-dl-us-salesforce-restricted-dev
irt-dl-us-salesforce-unrestricted
irt-dl-us-salesforce-unrestricted-dev
```
