// notificaciones.go
package entities

type Cita struct {
	CitaID           int32  `json:"citaId" gorm:"column:cita_id;primaryKey;autoIncrement"`
	NombrePaciente   string `json:"nombrePaciente" gorm:"column:nombre_paciente;not null"`
	ApellidoPaciente string `json:"apellidoPaciente" gorm:"column:apellido_paciente;not null"`
	NumeroContacto   string `json:"numeroContacto" gorm:"column:numero_contacto;not null"`
	AreaCita         string `json:"areaCita" gorm:"type:enum('Medicina General', 'Pediatría', 'Oftalmología', 'Dermatología', 'Cardiología');not null"`
	Fecha            string `json:"fecha" gorm:"column:fecha;not null"`
	Hora             string `json:"hora" gorm:"column:hora;not null"`
	Estado           string `json:"estado" gorm:"column:estado;not null;default:'pendiente'"`
}

// Setters
func (c *Cita) SetNombrePaciente(nombre string) {
	c.NombrePaciente = nombre
}

func (c *Cita) SetApellidoPaciente(apellido string) {
	c.ApellidoPaciente = apellido
}

func (c *Cita) SetNumeroContacto(numero string) {
	c.NumeroContacto = numero
}

func (c *Cita) SetAreaCita(area string) {
	c.AreaCita = area
}

func (c *Cita) SetFecha(fecha string) {
	c.Fecha = fecha
}

func (c *Cita) SetHora(hora string) {
	c.Hora = hora
}

func (c *Cita) SetEstado(estado string) {
	c.Estado = estado
}

// Getters
func (c *Cita) GetCitaID() int32 {
	return c.CitaID
}

func (c *Cita) GetNombrePaciente() string {
	return c.NombrePaciente
}

func (c *Cita) GetApellidoPaciente() string {
	return c.ApellidoPaciente
}

func (c *Cita) GetNumeroContacto() string {
	return c.NumeroContacto
}

func (c *Cita) GetAreaCita() string {
	return c.AreaCita
}

func (c *Cita) GetFecha() string {
	return c.Fecha
}

func (c *Cita) GetHora() string {
	return c.Hora
}

func (c *Cita) GetEstado() string {
	return c.Estado
}

func NewCita(nombrePaciente, apellidoPaciente, numeroContacto, areaCita, fecha, hora, estado string) *Cita {
	return &Cita{
		NombrePaciente:   nombrePaciente,
		ApellidoPaciente: apellidoPaciente,
		NumeroContacto:   numeroContacto,
		AreaCita:         areaCita,
		Fecha:            fecha,
		Hora:             hora,
		Estado:           estado,
	}
}
