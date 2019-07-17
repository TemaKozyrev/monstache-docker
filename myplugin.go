package main

import "github.com/rwynn/monstache/monstachemap"

func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	doc := input.Document
	tempDoc := map[string]interface{}{}
	if input.Collection == "forum_messages" {
		tempDoc["_id"] = doc["_id"]
		tempDoc["ownerId"] = doc["ownerId"]
		tempDoc["topicId"] = doc["topicId"]
		tempDoc["body"] = doc["body"]
		tempDoc["createdAt"] = doc["createdAt"]
	} else if input.Collection == "forum_topics" {
		tempDoc["_id"] = doc["_id"]
		tempDoc["title"] = doc["title"]
		tempDoc["createdAt"] = doc["createdAt"]
	} else if input.Collection == "forum_faq_articles" {
		tempDoc["_id"] = doc["_id"]
		tempDoc["topicId"] = doc["topicId"]
		tempDoc["titleId"] = doc["titleId"]
		tempDoc["title"] = doc["title"]
		tempDoc["preview"] = doc["preview"]
		tempDoc["body"] = doc["body"]
		tempDoc["createdAt"] = doc["createdAt"]
	} else if input.Collection == "forum_news" {
		tempDoc["_id"] = doc["_id"]
		tempDoc["title"] = doc["title"]
		tempDoc["body"] = doc["body"]
		tempDoc["createdAt"] = doc["createdAt"]
	}
	r := &monstachemap.MapperPluginOutput{Document: tempDoc}
	r.Drop = doc["isDeleted"].(bool)
	return r, nil
}
