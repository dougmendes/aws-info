package services

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListUsers(svc *iam.Client) {
	listUserOutput, err := svc.ListUsers(context.TODO(), &iam.ListUsersInput{})
	if err != nil {
		log.Fatalf("Unable to list users, %v", err)
	}

	fmt.Println("Users with their associated roles:")
	for _, user := range listUserOutput.Users {
		fmt.Printf("User: %s\n", *user.UserName)
		ListAttachedUserPolicies(svc, user.UserName)
		ListGroupsForUser(svc, user.UserName)
	}
}

func ListAttachedUserPolicies(svc *iam.Client, userName *string) {
	listAttachedUserPoliciesOutput, err := svc.ListAttachedUserPolicies(context.TODO(), &iam.ListAttachedUserPoliciesInput{
		UserName: userName,
	})
	if err != nil {
		log.Printf("Failed to list attached user policies for user %s, %v", *userName, err)
		return
	}

	if len(listAttachedUserPoliciesOutput.AttachedPolicies) == 0 {
		fmt.Println("  No attached policies")
		return
	}

	fmt.Println("  Attached Policies:")
	for _, policy := range listAttachedUserPoliciesOutput.AttachedPolicies {
		fmt.Printf("    %s\n", *policy.PolicyName)
	}
}

func ListGroupsForUser(svc *iam.Client, userName *string) {
	listGroupsForUserOutput, err := svc.ListGroupsForUser(context.TODO(), &iam.ListGroupsForUserInput{
		UserName: userName,
	})
	if err != nil {
		log.Printf("Failed to list groups for user %s, %v", *userName, err)
		return
	}

	if len(listGroupsForUserOutput.Groups) == 0 {
		fmt.Println("  No groups")
		return
	}

	fmt.Println("  Groups:")
	for _, group := range listGroupsForUserOutput.Groups {
		fmt.Printf("    %s\n", *group.GroupName)
		ListAttachedGroupPolicies(svc, group.GroupName)
	}
}

func ListAttachedGroupPolicies(svc *iam.Client, groupName *string) {
	listAttachedGroupPoliciesOutput, err := svc.ListAttachedGroupPolicies(context.TODO(), &iam.ListAttachedGroupPoliciesInput{
		GroupName: groupName,
	})
	if err != nil {
		log.Printf("Failed to list attached group policies for group %s, %v", *groupName, err)
		return
	}

	if len(listAttachedGroupPoliciesOutput.AttachedPolicies) == 0 {
		fmt.Println("    No attached group policies")
		return
	}

	fmt.Println("    Attached Group Policies:")
	for _, policy := range listAttachedGroupPoliciesOutput.AttachedPolicies {
		fmt.Printf("      %s\n", *policy.PolicyName)
	}
}
