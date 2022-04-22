package handler

import (
	"context"
	"place-srv/db"
	pb "place-srv/proto"
)

type PlaceServiceExtHandler struct {

}

func NewPlaceServiceExtHandler() *PlaceServiceExtHandler {
	return &PlaceServiceExtHandler{

	}
}

func (p *PlaceServiceExtHandler) HotCitiesByCinema(ctx context.Context, req *pb.HotCitiesByCinemaReq, rsp *pb.HotCitiesByCinemaRep) error {

	places,err := db.SelectPlaces()
	if err != nil {

		return err
	}
	placesPB := []*pb.PlaceData{}
	for _,place := range places {
		placePB := place.ToProtoDBHotPlayMovies()
		placesPB = append(placesPB, placePB)
	}
	rsp.P = placesPB
	return nil
}


