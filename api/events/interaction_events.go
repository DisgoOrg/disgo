package events

import (
	"github.com/DisgoOrg/disgo/api"
)

// GenericInteractionEvent generic api.Interaction event
type GenericInteractionEvent struct {
	GenericEvent
	Interaction *api.Interaction
}

// Respond replies to the api.Interaction with the provided api.InteractionResponse
func (e *GenericInteractionEvent) Respond(responseType api.InteractionResponseType, data interface{}) error {
	return e.Interaction.Respond(responseType, data)
}

// DeferReply replies to the api.CommandInteraction with api.InteractionResponseTypeDeferredChannelMessageWithSource and shows a loading state
func (e *GenericInteractionEvent) DeferReply(ephemeral bool) error {
	return e.Interaction.DeferReply(ephemeral)
}

// Reply replies to the api.Interaction with api.InteractionResponseTypeDeferredChannelMessageWithSource & api.MessageCreate
func (e *GenericInteractionEvent) Reply(messageCreate api.MessageCreate) error {
	return e.Interaction.Reply(messageCreate)
}

// EditOriginal edits the original api.InteractionResponse
func (e *GenericInteractionEvent) EditOriginal(messageUpdate api.MessageUpdate) (*api.Message, error) {
	return e.Interaction.EditOriginal(messageUpdate)
}

// DeleteOriginal deletes the original api.InteractionResponse
func (e *GenericInteractionEvent) DeleteOriginal() error {
	return e.Interaction.DeleteOriginal()
}

// SendFollowup used to send a followup api.MessageCreate to an api.Interaction
func (e *GenericInteractionEvent) SendFollowup(messageCreate api.MessageCreate) (*api.Message, error) {
	return e.Interaction.SendFollowup(messageCreate)
}

// EditFollowup used to edit a followup api.Message from an api.Interaction
func (e *GenericInteractionEvent) EditFollowup(messageID api.Snowflake, messageUpdate api.MessageUpdate) (*api.Message, error) {
	return e.Interaction.EditFollowup(messageID, messageUpdate)
}

// DeleteFollowup used to delete a followup api.Message from an api.Interaction
func (e *GenericInteractionEvent) DeleteFollowup(messageID api.Snowflake) error {
	return e.Interaction.DeleteFollowup(messageID)
}

// CommandEvent indicates that a slash api.Command was ran
type CommandEvent struct {
	GenericInteractionEvent
	CommandInteraction  *api.CommandInteraction
	CommandID           api.Snowflake
	CommandName         string
	SubCommandName      *string
	SubCommandGroupName *string
	Options             []*api.Option
}

// CommandPath returns the api.Command path
func (e CommandEvent) CommandPath() string {
	path := e.CommandName
	if e.SubCommandName != nil {
		path += "/" + *e.SubCommandName
	}
	if e.SubCommandGroupName != nil {
		path += "/" + *e.SubCommandGroupName
	}
	return path
}

// Option returns an Option by name
func (e CommandEvent) Option(name string) *api.Option {
	options := e.OptionN(name)
	if len(options) == 0 {
		return nil
	}
	return options[0]
}

// OptionN returns Option(s) by name
func (e CommandEvent) OptionN(name string) []*api.Option {
	options := make([]*api.Option, 0)
	for _, option := range e.Options {
		if option.Name == name {
			options = append(options, option)
		}
	}
	return options
}

// OptionsT returns Option(s) by api.CommandOptionType
func (e CommandEvent) OptionsT(optionType api.CommandOptionType) []*api.Option {
	options := make([]*api.Option, 0)
	for _, option := range e.Options {
		if option.Type == optionType {
			options = append(options, option)
		}
	}
	return options
}

// ButtonClickEvent indicates that a api.Button was clicked
type ButtonClickEvent struct {
	GenericInteractionEvent
	ButtonInteraction *api.ButtonInteraction
}

// DeferEdit replies to the api.ButtonInteraction with api.InteractionResponseTypeDeferredUpdateMessage and cancels the loading state
func (e *ButtonClickEvent) DeferEdit() error {
	return e.ButtonInteraction.DeferEdit()
}

// Edit replies to the api.ButtonInteraction with api.InteractionResponseTypeUpdateMessage & api.MessageCreate which edits the original api.Message
func (e *ButtonClickEvent) Edit(messageCreate api.MessageCreate) error {
	return e.ButtonInteraction.Edit(messageCreate)
}

// CustomID returns the customID from the called api.Button
func (e *ButtonClickEvent) CustomID() string {
	return e.ButtonInteraction.Data.CustomID
}

// ComponentType returns the api.ComponentType from the called api.Button
func (e *ButtonClickEvent) ComponentType() string {
	return e.ButtonInteraction.Data.CustomID
}

// Message returns the api.Message the api.Button is called from
func (e *ButtonClickEvent) Message() *api.Message {
	return e.ButtonInteraction.Message
}
