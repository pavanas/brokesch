package grifts

import (
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
	"github.com/pavanas/brokestest/models"
	//"time"
	"math/rand"
	"github.com/satori/go.uuid"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		
		// if err := models.DB.TruncateAll(); err != nil {
	 //      return errors.WithStack(err)
	 //    }

	    //var m_ids[3]string;
	    r_ids := [20]string{
	    	 "70ef2c1f-9ece-42b4-bda0-dad7852a105e",
			 "6a3aa578-3078-4502-abc9-31b26ebe1332",
			 "c9526031-ea46-4530-8c02-02de7b81f5cb",
			 "94483b97-0b9e-457e-ba99-bb9c95fa1349",
			 "bb1fb18e-ae61-4114-8cf9-a889b636e9a3",
			 "aa08fa49-386c-46ea-9638-ddb716b6dfb4",
			 "908f7cfc-cf84-4361-a24b-d71ce98a5f15",
			 "becc1a5b-7917-4292-a2c0-00e774f6590e",
			 "c68520e4-67e0-4936-98f1-836226b62885",
			 "de16738e-964f-4d1b-bfb2-e3cde8d778a5",
			 "16bed742-9c9f-4796-8967-0228ee4939f1",
			 "b6abb398-1a1e-4676-ad19-b4addcc34e80",
			 "2219600c-9fc5-4a3e-97f4-b3c4b21f8c34",
			 "aa4378a3-601f-465c-829b-d83aca8e0f9a",
			 "40b77d08-c51e-4471-85c2-f5af0b0ae76b",
			 "3e875408-6047-42d0-9390-e6718c2699fc",
			 "4ea6548b-1222-457f-93d1-79bbb24b5840",
			 "6a5fa206-16b3-4c29-a16f-fad435c51c39",
			 "2c18a60b-65b0-45c7-9164-3918c204ea94",
			 "30732aef-40c1-499a-b7d8-fab3f3376022",
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

	    // meetings := [3]string{"5629d57b-4e1d-4a07-b095-9eaa668510e0","704fc452-7d78-4c73-86e7-d71c918ed5ed","d8f61ca0-6dd8-4097-8a6c-e02c4f290660"}
	   	// // Seeding Races
	   	// for i :=0; i<20; i++ {
	   	// 	uuid_val := meetings[rand.Intn(3)]
	   	// 	//uuid_id := r_ids[rand.Intn(20)]
	   	// 	ends_at_t := time.Now().Local().Add(time.Duration(45*i) * time.Minute)
	   	// 	race := &models.Race{
	   	// 		Number: i+1,
	   	// 		MeetingID: uuid.FromStringOrNil(uuid_val),
	   	// 		EndsAt: ends_at_t,
	   	// 	}
	   	// 	if err := models.DB.Create(race); err != nil {
		   //    return errors.WithStack(err)
		   //  }
	   	// }

	   	//Seeding competitors
	   	for i := 0; i<150; i++ {
	   		uuid_val := r_ids[rand.Intn(20)]
	   		competitor := &models.Competitor {
	   			Position: rand.Intn(8) + 1,
	   			Number: rand.Intn(999) + 1,
	   			Race_id: uuid.FromStringOrNil(uuid_val),
	   		}
	   		if err := models.DB.Create(competitor); err != nil {
		      return errors.WithStack(err)
		    }
	   	}

		return nil
	})

})
