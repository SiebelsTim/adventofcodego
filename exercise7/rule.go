package exercise7

import (
	"errors"
	"strconv"
	"strings"
)

type rule struct {
	outerBag      bag
	containedBags []containedBag
}

func parseRule(line string) (rule, error) {
	words := strings.SplitN(line, " ", 3)
	outerBag := NewBag(words[0], words[1])

	var containedBags []containedBag

	rest := line[strings.Index(line, "contain ")+len("contain "):]
	if rest != "no other bags." {
		innerBags := strings.Split(rest, ", ")
		for _, innerBag := range innerBags {
			parts := strings.Split(innerBag, " ")
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				return rule{}, errors.New("could not parse " + parts[0] + " as int")
			}
			containedBags = append(containedBags, containedBag{
				NewBag(parts[1], parts[2]),
				count,
			})
		}
	}

	return rule{
		outerBag:      outerBag,
		containedBags: containedBags,
	}, nil
}
