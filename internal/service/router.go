package service

import (
	"github.com/go-chi/chi"
	"github.com/rarimo/registration-relayer/internal/service/contract"
	"github.com/rarimo/registration-relayer/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	registrationABI, err := contract.RegistrationMetaData.GetAbi()
	if err != nil {
		panic(errors.Wrap(err, "failed to get vote verifier ABI"))
	}

	method, ok := registrationABI.Methods["register"]
	if !ok {
		panic(errors.New("register method not found"))
	}

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxRelayerConfig(s.cfg.RelayerConfig()),
			handlers.CtxRegisterMethod(&method),
		),
	)
	r.Route("/integrations/registration-relayer", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/register", handlers.Registration)
		})
	})

	return r
}
