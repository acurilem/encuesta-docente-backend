package models

import (
	"time"
)

type AdminPostUmagPsu struct {
	TipoIden            string    `gorm:"column:tipo_iden" json:"tipoIden"`
	RutIns              string    `gorm:"column:rut_ins" json:"rutIns"`
	AnoProc             int       `gorm:"column:ano_proc" json:"anoProc"`
	Paterno             string    `gorm:"column:paterno" json:"paterno"`
	Materno             string    `gorm:"column:materno" json:"materno"`
	Nombres             string    `gorm:"column:nombres" json:"nombres"`
	Nacionalidad        int       `gorm:"column:nacionalidad" json:"nacionalidad"`
	Sexo                string    `gorm:"column:sexo" json:"sexo"`
	Preferencia         int       `gorm:"column:preferencia" json:"preferencia"`
	CodCar              string    `gorm:"column:cod_car" json:"codCar"`
	EstPost             int       `gorm:"column:est_post" json:"estPost"`
	PuntPond            int       `gorm:"column:punt_pond" json:"puntPond"`
	Lugar               int       `gorm:"column:lugar" json:"lugar"`
	SitEgrEduc          int       `gorm:"column:sit_egr_educ" json:"sitEgrEduc"`
	LocalEduc           int       `gorm:"column:local_educ" json:"localEduc"`
	UnidEduc            int       `gorm:"column:unid_educ" json:"unidEduc"`
	CodRama             string    `gorm:"column:cod_rama" json:"codRama"`
	GrupDepCol          int       `gorm:"column:grup_dep_col" json:"grupDepCol"`
	CodRegion           int       `gorm:"column:cod_region" json:"codRegion"`
	CodProv             int       `gorm:"column:cod_prov" json:"codProv"`
	AnoEgre             int       `gorm:"column:ano_egre" json:"anoEgre"`
	PromNotas           int       `gorm:"column:prom_notas" json:"promNotas"`
	PuntNem             int       `gorm:"column:punt_nem" json:"puntNem"`
	Lyc                 int       `gorm:"column:lyc" json:"lyc"`
	Mat                 int       `gorm:"column:mat" json:"mat"`
	Mat2                int       `gorm:"column:mat2" json:"mat2"`
	Hycs                int       `gorm:"column:hycs" json:"hycs"`
	Cs                  int       `gorm:"column:cs" json:"cs"`
	PromPsu             int       `gorm:"column:prom_psu" json:"promPsu"`
	FecNac              time.Time `gorm:"column:fec_nac" json:"fecNac"`
	IngBrutFam          int       `gorm:"column:ing_brut_fam" json:"ingBrutFam"`
	Salud               int       `gorm:"column:salud" json:"salud"`
	RutPadre            string    `gorm:"column:rut_padre" json:"rutPadre"`
	RutMadre            string    `gorm:"column:rut_madre" json:"rutMadre"`
	Domicilio           string    `gorm:"column:domicilio" json:"domicilio"`
	EmailIns            string    `gorm:"column:email_ins" json:"emailIns"`
	FolioTp             int       `gorm:"column:folio_tp" json:"folioTp"`
	Repostula           string    `gorm:"column:repostula" json:"repostula"`
	TeleOri             int       `gorm:"column:tele_ori" json:"teleOri"`
	Ciudad              string    `gorm:"column:ciudad" json:"ciudad"`
	CodRegionDom        int       `gorm:"column:cod_region_dom" json:"codRegionDom"`
	Bea                 string    `gorm:"column:bea" json:"bea"`
	ModuloCiencias      string    `gorm:"column:modulo_ciencias" json:"moduloCiencias"`
	PondAnoAcad         int       `gorm:"column:pond_ano_acad" json:"pondAnoAcad"`
	IdAdminPost         int       `gorm:"column:id_admin_post" json:"idAdminPost"`
	DireccionActual     string    `gorm:"column:direccion_actual" json:"direccionActual"`
	TelefonoActual      string    `gorm:"column:telefono_actual" json:"telefonoActual"`
	PuntRanking         int       `gorm:"column:punt_ranking" json:"puntRanking"`
	RutOriginal         string    `gorm:"column:rut_original" json:"rutOriginal"`
	PublicaNombreSocial string    `gorm:"column:publica_nombre_social" json:"publicaNombreSocial"`
	NombreSocial        string    `gorm:"column:nombre_social" json:"nombreSocial"`
	PuebloEtniaDatos    string    `gorm:"column:pueblo_etnia_datos" json:"puebloEtniaDatos"`
	CodigoEtnia         int       `gorm:"column:codigo_etnia" json:"codigoEtnia"`
	P50                 string    `gorm:"column:p_50" json:"p50"`
	P60                 string    `gorm:"column:p_60" json:"p60"`
}

func (AdminPostUmagPsu) TableName() string {
	return "admin_post_umag_psu"
}
