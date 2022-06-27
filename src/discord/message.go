package discord

import (
	"TsunoKento/AWS-server-management-BOT/start"
	"TsunoKento/AWS-server-management-BOT/stop"
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/bwmarrin/discordgo"
)

var (
	awsconfig aws.Config
	client    *ec2.Client
)

func init() {
	var err error
	awsconfig, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client = ec2.NewFromConfig(awsconfig)
}

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	st, _ := s.Channel(m.ChannelID)

	if st.Name == "arkサーバー" {
		if m.Content == "開始" {
			input := &ec2.StartInstancesInput{
				InstanceIds: []string{
					os.Getenv("AWS_INSTANCE_ID"),
				},
				DryRun: aws.Bool(true),
			}

			if _, err := start.StartInstance(context.TODO(), client, input); err != nil {
				log.Fatal(err)
			}
			s.ChannelMessageSend(m.ChannelID, "サーバーを開始します")
		} else if m.Content == "終了" {
			input := &ec2.StopInstancesInput{
				InstanceIds: []string{
					os.Getenv("AWS_INSTANCE_ID"),
				},
				DryRun: aws.Bool(true),
			}

			if _, err := stop.StopInstance(context.TODO(), client, input); err != nil {
				log.Fatal(err)
			}
			s.ChannelMessageSend(m.ChannelID, "サーバーを終了しました")
		}
	}
}
