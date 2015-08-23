package adapters

import (
	"log"
	"testing"
)

func TestGetDate(t *testing.T) {
	rp := NewReadovkaParser()
	/*var dates = map[string]time.Time{
		"21 ИЮНЬ 2015 09:46": time.Time{},
	}*/
	date, err := rp.getDate("21 ИЮНЬ 2015 09:46")
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(date)
	}

}

func TestGetDateUpdated(t *testing.T) {
	rp := NewReadovkaParser()
	var html = `<span class="item-date">
		  20 Июнь 2015 19:20				
						<!-- Item date modified -->
			<span class="itemDateModified" style="font-size:0.8em;color:#777">
				<span style="text-transform:lowercase">(Последнее изменение</span> Суббота, 20 Июнь 2015 20:24)
			</span>
			
				
                  
		</span>`
	date, err := rp.getDate(html)
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(date)
	}
}
