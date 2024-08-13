package client

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type db interface {
	StoreName(name string) (string, error)
	GetName(uuid string) (string, error)
}

type logger interface {
	Debug(message string)
	Error(err error)
}

type client struct {
	logger     logger
	db         db
	numPlayers int
	playerKeys []string
}

func NewClient(logger logger, db db) (*client, error) {
	if db == nil {
		return nil, errors.New("tried to create new client, db was nil")
	}
	return &client{
		logger:     logger,
		db:         db,
		playerKeys: []string{},
	}, nil
}

func (c *client) AddPlayers() error {
	if c.db == nil {
		return errors.New("can't add player, db was nil")
	}

	key, err := c.db.StoreName("Chip")
	if err != nil {
		return err
	}
	c.playerKeys = append(c.playerKeys, key)

	key, err = c.db.StoreName("Dale")
	if err != nil {
		return err
	}
	c.playerKeys = append(c.playerKeys, key)

	key, err = c.db.StoreName("Monterey Jack")
	if err != nil {
		return err
	}
	c.playerKeys = append(c.playerKeys, key)

	key, err = c.db.StoreName("Gadget Hackwrench")
	if err != nil {
		return err
	}
	c.playerKeys = append(c.playerKeys, key)

	key, err = c.db.StoreName("Zipper Fly")
	if err != nil {
		return err
	}
	c.playerKeys = append(c.playerKeys, key)

	key, err = c.db.StoreName("Fat Cat")
	if err != nil {
		return err
	}
	c.playerKeys = append(c.playerKeys, key)

	key, err = c.db.StoreName("Rat Capone")
	if err != nil {
		return err
	}
	c.playerKeys = append(c.playerKeys, key)

	c.numPlayers = len(c.playerKeys)
	return nil
}

func (c *client) PlayRounds(numRounds int) error {
	if c.db == nil {
		return errors.New("tried to play a round, db was nil")
	}

	if c.numPlayers < 2 {
		return errors.New("not enough players to play a round")
	}

	i := 0
	for i < numRounds {
		c.logger.Debug(fmt.Sprintf("-= Playing Round %d =-", i+1))
		p1 := rand.Intn(c.numPlayers - 1)
		p2 := rand.Intn(c.numPlayers - 1)
		if p1 == p2 {
			c.logger.Debug("Round ended in a draw!")
			time.Sleep(time.Second * 2)
			continue
		}

		player1 := c.playerKeys[p1]
		p1Name, err := c.db.GetName(player1)
		if err != nil {
			c.logger.Error(err)
		}

		player2 := c.playerKeys[p2]
		p2Name, err := c.db.GetName(player2)
		if err != nil {
			c.logger.Error(err)
		}

		c.logger.Debug(fmt.Sprintf("PLayer %s attacked Player %s !", p1Name, p2Name))
		time.Sleep(time.Second * 2)

		winner := rand.Intn(2)
		if winner == 0 {
			c.logger.Debug(fmt.Sprintf("%s won!", p1Name))
		} else {
			c.logger.Debug(fmt.Sprintf("%s won!", p2Name))
		}
		c.logger.Debug("-= ROUND FINISHED =- \n")

		i++
	}
	return nil
}
