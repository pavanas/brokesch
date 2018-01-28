package grifts

import (
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
	"github.com/pavanas/brokestest/models"
	"time"
	"math/rand"
	"github.com/satori/go.uuid"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		
		if err := models.DB.TruncateAll(); err != nil {
	      return errors.WithStack(err)
	    }

	    // Seeding meetings
	    /*meetings := [3]string{"thoroubred", "harness", "greyhound"}
	    for i := 0; i < 3; i++ {
	    	meeting := &models.Meeting{
		      Type: meetings[i],
		    }
		    if err := models.DB.Create(meeting); err != nil {
		      return errors.WithStack(err)
		    }
	    }*/

	    meetings := [3]string{"021DB630-FF8E-42AA-8ABD-62ECEF3FB275","54DE7778-8BDD-4AE3-8223-4291B6B671BA","E47E4F9E-AF01-4769-8CD5-ED4889DCC8AE"}
	   	// Seeding Races
	   	for i :=0; i<20; i++ {
	   		uuid_val := meetings[rand.Intn(3)]
	   		ends_at_t := time.Now().Local().Add(time.Duration(15*i) * time.Minute)
	   		race := &models.Race{
	   			Number: i+1,
	   			MeetingID: uuid.FromStringOrNil(uuid_val),
	   			EndsAt: ends_at_t,
	   		}
	   		if err := models.DB.Create(race); err != nil {
		      return errors.WithStack(err)
		    }
	   	}

	   	// Seeding competitors
	   	/*for i := 0; i<100; i++ {
	   		competitor := &models.Competitor {
	   			Position: rand.Intn(8) + 1,
	   			Number: rand.Intn(999) + 1,
	   		}
	   		if err := models.DB.Create(competitor); err != nil {
		      return errors.WithStack(err)
		    }
	   	}*/

		return nil
	})

})
