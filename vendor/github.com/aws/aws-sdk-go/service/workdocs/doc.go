// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package workdocs provides the client and types for making API
// requests to Amazon WorkDocs.
//
// The WorkDocs API is designed for the following use cases:
//
//    * File Migration: File migration applications are supported for users
//    who want to migrate their files from an on-premise or off-premise file
//    system or service. Users can insert files into a user directory structure,
//    as well as allow for basic metadata changes, such as modifications to
//    the permissions of files.
//
//    * Security: Support security applications are supported for users who
//    have additional security needs, such as anti-virus or data loss prevention.
//    The APIs, in conjunction with Amazon CloudTrail, allow these applications
//    to detect when changes occur in Amazon WorkDocs, so the application can
//    take the necessary actions and replace the target file. The application
//    can also choose to email the user if the target file violates the policy.
//
//    * eDiscovery/Analytics: General administrative applications are supported,
//    such as eDiscovery and analytics. These applications can choose to mimic
//    and/or record the actions in an Amazon WorkDocs site, in conjunction with
//    Amazon CloudTrails, to replicate data for eDiscovery, backup, or analytical
//    applications.
//
// All Amazon WorkDocs APIs are Amazon authenticated, certificate-signed APIs.
// They not only require the use of the AWS SDK, but also allow for the exclusive
// use of IAM users and roles to help facilitate access, trust, and permission
// policies. By creating a role and allowing an IAM user to access the Amazon
// WorkDocs site, the IAM user gains full administrative visibility into the
// entire Amazon WorkDocs site (or as set in the IAM policy). This includes,
// but is not limited to, the ability to modify file permissions and upload
// any file to any user. This allows developers to perform the three use cases
// above, as well as give users the ability to grant access on a selective basis
// using the IAM model.
//
// See https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01 for more information on this service.
//
// See workdocs package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/workdocs/
//
// Using the Client
//
// To use the client for Amazon WorkDocs you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := workdocs.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon WorkDocs client WorkDocs for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/workdocs/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.AbortDocumentVersionUpload(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("AbortDocumentVersionUpload result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.AbortDocumentVersionUploadWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package workdocs
