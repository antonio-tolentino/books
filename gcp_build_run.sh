

gcloud builds submit --tag gcr.io/arquitetura-207620/books:v1.0.0


gcloud run deploy --image gcr.io/arquitetura-207620/books:v1.0.0



## token for test
gcloud auth print-identity-token