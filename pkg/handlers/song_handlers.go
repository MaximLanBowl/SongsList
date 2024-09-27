package handlers

import (
	"SongsList/pkg/models"
	"SongsList/pkg/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get list of songs
// @Description Get all songs
// @Tags songs
// @Success 200 {array} models.Song
// @Router /songs [get]
func GetSongs(ctx *gin.Context) {
	var songs []models.Song
	err := repository.DB.Select(&songs, "SELECT * FROM songs")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, songs)
}

// @Summary Create song by ID
// @Description Create song by its ID
// @Tags songs
// @Accept json
// @Produce json
// @Success 200 {object} models.Song
// @Failure 404 {object} string
// @Router /songs [post]
func CreateSong(ctx *gin.Context) {
	var input models.Song
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := `INSERT INTO songs (group_name, song_name, release_date, text, link) 
	VALUES (:group_name, :song_name, :release_date, :text, :link) RETURNING id`

	var id string
	_, err := repository.DB.NamedExec(query, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	input.ID = id
	ctx.JSON(http.StatusOK, input)

}

// @Summary Get song by ID
// @Description Get song by its ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} models.Song
// @Failure 404 {object} string
// @Router /songs/{id} [get]
func GetSongByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var song models.Song
	err := repository.DB.Get(&song, "SELECT * FROM songs WHERE id=$1", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, song)
}

// @Summary Delete song by ID
// @Description Delete song by its ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} models.Song
// @Failure 404 {object} string
// @Router /songs/{id} [post]
func DeleteSong(ctx *gin.Context) {
    id := ctx.Param("id")
    _, err := repository.DB.Exec("DELETE FROM songs WHERE id=$1", id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"status": "Song deleted successfully"})
}

// @Summary Put song by ID
// @Description Update song by its ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} models.Song
// @Failure 404 {object} string
// @Router /songs/{id} [post]
func UpdateSong(ctx *gin.Context) {
    id := ctx.Param("id")
    var input models.Song
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	input.ID = id
    query := `UPDATE songs SET group_name=:group_name, song_name=:song_name, release_date=:release_date, text=:text, link=:link WHERE id=:id`
    _, err := repository.DB.NamedExec(query, input)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"status": "Song updated successfully"})
}