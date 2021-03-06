package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// projeto represents data about a record projeto.
type projeto struct {
    ID_Projeto     string  `json:"id"`
    Title  string  `json:"title"`
    Description string  `json:"Description"`
    IDequipe string `json:"equipe"`
    IDtarefa string `json:"tarefas"`
}

type pessoa struct {
    ID_Pessoa     string  `json:"id"`
    Nome  string  `json:"nome"`
    Profissao string  `json:"profissao"`
    ID_Tarefa string `json:"id_tarefa"`
}

type equipe struct {
    Nome string `json:"nome"`
    ID_Equipe string   `json:"id"`
    IDMembers string `json:"idmembers"`
}

type tarefa struct {
    ID_Tarefa string `json:"id"`
    Nome string `json:"nome"`
    Description string `json:"description"`
    ID_Project string `json:"ID_Projeto"`
    ID_Equipe string `json:"ID_Equipe"`
    Tempo string `json:"tempo"`
}

// projetos slice to seed record projeto data.
var projetos = []projeto{
    {ID_Projeto: "1", Title: "Central de Relacionamento", Description: "Sugestões", IDequipe: "1", IDtarefa: "1, 2, 4"},
    {ID_Projeto: "2", Title: "Jeru", Description: "talvez de certo", IDequipe: "2", IDtarefa: "2"},
    {ID_Projeto: "3", Title: "Sarah Vaughan and Clifford Brown", Description: "talvez de certo", IDequipe: "3", IDtarefa: "4,5"},
}

var pessoas = []pessoa{
    {ID_Pessoa: "1", Nome: "Bruno", Profissao: "45", ID_Tarefa: "2"},
    {ID_Pessoa: "2", Nome: "Pedro", Profissao: "12", ID_Tarefa: "1"},
    {ID_Pessoa: "3", Nome: "Caio",  Profissao: "13", ID_Tarefa: "3"},
}

var equipes = []equipe{
    {ID_Equipe: "1", Nome: "Komanda", IDMembers: "3, 2, 1"},
    {ID_Equipe: "2", Nome: "DevsCariri", IDMembers: ""},
    {ID_Equipe: "3", Nome: "Kariri Inovação", IDMembers: ""},
}

var tarefas = []tarefa {
    {ID_Tarefa: "1", Nome: "Criação de API REST", Description: "Utilizar GO LANG com Gin", ID_Project: "1", ID_Equipe: "", Tempo: ""},
    {ID_Tarefa: "2", Nome: "Teste", Description: "Apenas Teste", ID_Project: "1", ID_Equipe: "", Tempo: ""},
    {ID_Tarefa: "3", Nome: "Teste", Description: "Apenas Teste", ID_Project: "1", ID_Equipe: "", Tempo: ""},
    {ID_Tarefa: "4", Nome: "Teste", Description: "Apenas Teste", ID_Project: "1", ID_Equipe: "", Tempo: ""},
    {ID_Tarefa: "5", Nome: "Teste", Description: "Apenas Teste", ID_Project: "1", ID_Equipe: "", Tempo: ""},
}

func main() {
    router := gin.Default()
    router.GET("/projetos", getprojetos)
    router.GET("/projetos/:id", getprojetoByID)
    router.GET("/projetos/:id/tarefas", getTarefasByProject)
    router.POST("/projetos", postprojetos)
    router.PUT("/projetos/:id", editProjetoById)
    router.DELETE("/projetos/:id", deleteProjetoById)
    router.GET("/projetos/equipes/:id", getEquipeByID)
    router.GET("/projetos/equipes/:id/members", getMembersInEquipeByID)
    router.POST("/projetos/:id/tarefas", postTarefaProjeto)


    router.GET("/tarefas", getTarefas)
    router.GET("/tarefas/:id", getTarefaByID)
    router.POST("/tarefas", postTarefas)
    router.PUT("/tarefas/:id", editTarefaById)
    router.DELETE("/tarefas/:id", deleteTarefaById)

    router.GET("/equipes", getEquipes)
    router.GET("/equipes/:id", getEquipeByID)
    router.GET("/equipes/member/:id", getMemberByID)
    router.POST("/equipes", postEquipes)
    router.DELETE("/equipes/:id", deleteEquipeById)
    router.PUT("/equipes/:id", updateEquipeById)

    router.GET("/pessoas", getPessoas)
    router.GET("/pessoas/:id", getpessoaByID)
    //router.GET("pessoas/:id/tarefas", getPessoaByIDTarefa)
    router.POST("/pessoas", postpessoas)
    router.DELETE("/pessoas/:id", deletePessoaById)
    router.PUT("/pessoas/:id", updatePessoaById)

    router.Run("localhost:2828")
}

// getprojetos/Pessoas/Equipes responds with the list of all projetos as JSON.
func getprojetos(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, projetos)
}

func getPessoas(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, pessoas)
}

func getEquipes(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, equipes)
}

func getTarefas(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, tarefas)
}

// postprojetos adds an projeto from JSON received in the request body.
func postprojetos(c *gin.Context) {
    var newprojeto projeto

    // Call BindJSON to bind the received JSON to
    // newprojeto.
    if err := c.BindJSON(&newprojeto); err != nil {
        return
    }

    // Add the new projeto to the slice.
    projetos = append(projetos, newprojeto)
    c.IndentedJSON(http.StatusCreated, newprojeto)
}

// getprojetoByID locates the projeto whose ID value matches the id

func getprojetoByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range projetos {
        if a.ID_Projeto == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "projeto not found"})
}

    // Delete a project from the list of projects by ID
func deleteProjetoById(c *gin.Context) {
    id := c.Param("id")
    for i, a := range projetos {
        if a.ID_Projeto == id {
            projetos = append(projetos[:i], projetos[i+1:]... )
            return
        }
    }
}

    // Edit a project from the project list by ID
func editProjetoById(c *gin.Context) {
    id := c.Param("id")
    for i := range projetos {
        if projetos[i].ID_Projeto == id {
        c.BindJSON(&projetos[i])
        c.IndentedJSON(http.StatusOK,projetos[i])
        return
        }
    }
}

func postpessoas(c *gin.Context) {
    var newpessoa pessoa

    // Call BindJSON to bind the received JSON to
    // newpessoa.
    if err := c.BindJSON(&newpessoa); err != nil {
        return
    }

    // Add the new pessoa to the slice.
    pessoas = append(pessoas, newpessoa)
    c.IndentedJSON(http.StatusCreated, newpessoa)
}

func postEquipes(c *gin.Context) {
    var newequipe equipe
    // Call BindJSON to bind the received JSON to newpessoa
    if err := c.BindJSON(&newequipe); err != nil {
        return
    }
    // Add the new pessoa to the slice.
    equipes = append(equipes, newequipe)
    c.IndentedJSON(http.StatusCreated, newequipe)
}

func deleteEquipeById(c *gin.Context) {
    id := c.Param("id")
    for i, a := range equipes {
        if a.ID_Equipe == id {
            equipes = append(equipes[:i], equipes[i+1:]... )
            return
        }
    }
}

func updateEquipeById(c *gin.Context) {
    id := c.Param("id")
    for i := range equipes {
        if equipes[i].ID_Equipe == id {
        c.BindJSON(&equipes[i])
        c.IndentedJSON(http.StatusOK,equipes[i])
        return
        }
    }
}

func getpessoaByID(c *gin.Context) {
    id := c.Param("id")

    /* Loop through the list of pessoas, looking for
     an pessoa whose ID value matches the parameter.*/
    for _, a := range pessoas {
        if a.ID_Pessoa == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pessoa not found"})
}

func getMemberByID(c *gin.Context) {
    id := c.Param("id")

    /* Loop through the list of pessoas, looking for
     an pessoa whose ID value matches the parameter.*/
    for _, a := range pessoas {
        if a.ID_Pessoa == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Member not found"})
}

func getEquipeByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range equipes {
        if a.ID_Equipe == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Equipe not found"})
}      

func getMembersInEquipeByID(c *gin.Context) {
    id := c.Param("id")
    count := 0
    for _, a := range equipes {
        for _, i := range pessoas {
            if a.ID_Equipe == id {
            c.IndentedJSON(http.StatusOK, i)
            
            count += 1
        }
    }
    if count > 0 {
        return
    } else {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Equipe not found"})
    }
        }
        
}


    // Delete a person from the list of people by Id
func deletePessoaById(c *gin.Context) {
    id := c.Param("id")
    for i, a := range pessoas {
        if a.ID_Pessoa == id {
            pessoas = append(pessoas[:i], pessoas[i+1:]... )
            return
        }
    }
}
    // edit a person from a list of people via id
func updatePessoaById(c *gin.Context) {
    id := c.Param("id")
    for i := range pessoas {
        if pessoas[i].ID_Pessoa == id {
        c.BindJSON(&pessoas[i])
        c.IndentedJSON(http.StatusOK,pessoas[i])
        return
        }
    }
}

func getTarefaByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range tarefas {
        if a.ID_Tarefa == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Tarefa not found"})
}

func postTarefas(c *gin.Context) {
    var newtarefa tarefa

    // Call BindJSON to bind the received JSON to
    // newpessoa.
    if err := c.BindJSON(&newtarefa); err != nil {
        return
    }

    // Add the new pessoa to the slice.
    tarefas = append(tarefas, newtarefa)
    c.IndentedJSON(http.StatusCreated, newtarefa)
}

func editTarefaById(c *gin.Context) {
    id := c.Param("id")
    for i := range tarefas {
        if tarefas[i].ID_Tarefa == id {
        c.BindJSON(&tarefas[i])
        c.IndentedJSON(http.StatusOK,tarefas[i])
        return
        }
    }
}
// // Delete a tarefa from the list of tarefas by Id

func deleteTarefaById(c *gin.Context) {
    id := c.Param("id")
    for i, a := range tarefas {
        if a.ID_Tarefa == id {
            tarefas = append(tarefas[:i], tarefas[i+1:]... )
            return
        }
    }
}

func getTarefasByProject(c *gin.Context) {
	id := c.Param("id")
	count := 0
	for _, a := range tarefas {
		if a.ID_Project == id {
			c.IndentedJSON(http.StatusOK, a)
			count+=1
		}
	}
	if(count > 0){
		return
	} else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tarefa not found"})
	}
}

func postTarefaProjeto(c *gin.Context){
    id := c.Param("id")
	count := 0
    for _, a := range projetos {
        if a.ID_Projeto == id {
                var newtarefa tarefa
                // Call BindJSON to bind the received JSON to
                // newpessoa.
                if err := c.BindJSON(&newtarefa); err != nil {
                    return
                }
    
                // Add the new pessoa to the slice.
                tarefas = append(tarefas, newtarefa)
                c.IndentedJSON(http.StatusCreated, newtarefa)
                return
        }  
    }
	if(count > 0){
		return
	} else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tarefa not found"})
	}
}
/*
func getPessoaByIDTarefa2(c *gin.Context){
    id := c.Param("id")
	count := 0
    var p_idTarefa pessoas
	for _, a := range pessoas {
		if a.ID_Pessoa == id {
			c.IndentedJSON(http.StatusOK, a)
            //idpessoa = a.ID_Pessoa
			count+=1
            p_idtarefa = a.ID_Tarefa 

            for _, b := range tarefas {
                if b.ID_Tarefa == p_idTarefa {
                    c.IndentedJSON(http.StatusOK, b)
                    count+=1
                }
		}
        
	}
	if(count > 2){
		return
	} else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tarefa not found"})
	}
}
*/
