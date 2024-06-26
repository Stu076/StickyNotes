package note

import (
	"janstupica/StickyNotes/internal/app/models"
	"janstupica/StickyNotes/internal/app/note"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	log     *zerolog.Logger
	useCase note.UseCase
}

func New(log *zerolog.Logger, useCase note.UseCase) *Handler {
	return &Handler{
		log:     log,
		useCase: useCase,
	}
}

// @BasePath		/api/v1

// @Summary Create new Sticky Note
// @Schemes
// @Description Creates a new sticky note.
// @Tags note
// @Accept json
// @Produce json
// @Param   note body models.Note true "user password"
// @Success 200 {object} models.Note
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /note/add [post]
func (h *Handler) Create(ctx *gin.Context) {
	input := new(models.Note)

	if err := ctx.BindJSON(input); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	result, err := h.useCase.Create(input)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, *result)
}

// @BasePath		/api/v1

// @Summary Update existing Sticky Note
// @Schemes
// @Description Updates an existing sticky note.
// @Tags note
// @Accept json
// @Produce json
// @Param   note body models.Note true "user password"
// @Success 200 {object} models.Note
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /note/update [put]
func (h *Handler) Update(ctx *gin.Context) {
	note := new(models.Note)

	if err := ctx.BindJSON(note); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	result, err := h.useCase.Update(note)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, *result)
}

// @BasePath		/api/v1

// @Summary Delete existing Sticky Note
// @Schemes
// @Description Deletes an existing sticky note.
// @Tags note
// @Param id path int true "Note ID"
// @Success 200 {string} http.StatusOK
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /note/delete/{id} [delete]
func (h *Handler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if id == 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "missing ID"})
		log.Error().Msg("Missing ID for note.")
		return
	}

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	if err := h.useCase.Delete(id); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"Message": "OK"})
}

// @BasePath		/api/v1

// @Summary Get Sticky Note
// @Schemes
// @Description Fetches a sticky note.
// @Tags note
// @Param id path int true "Note ID"
// @Produce json
// @Success 200 {object} models.Note
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /note/get/{id} [get]
func (h *Handler) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if id == 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "missing ID"})
		log.Error().Msg("Missing id for note.")
		return
	}

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	result, err := h.useCase.Get(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, *result)
}

// @BasePath		/api/v1

// @Summary Get all Sticky Notes for User
// @Schemes
// @Description Fetches all sticky notes for user.
// @Tags note
// @Param userId path int true "User ID"
// @Produce json
// @Success 200 {object} models.Note
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /note/get/all/{userId} [get]
func (h *Handler) GetAll(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))

	if userId == 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "missing UserID"})
		log.Error().Msg("Missing userId for note.")
		return
	}

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	results, err := h.useCase.GetAll(userId)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		log.Error().Msg(err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, *results)
}
