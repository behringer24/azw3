package azw3

// marcRelatorCodes is the set of valid MARC relator codes accepted by
// AddContributor, sourced from the Library of Congress's MARC relator
// term vocabulary (https://id.loc.gov/vocabulary/relators.json).
var marcRelatorCodes = map[string]bool{
	"abr": true, // Abridger
	"acp": true, // Art copyist
	"act": true, // Actor
	"adi": true, // Art director
	"adp": true, // Adapter
	"aft": true, // Author of afterword, colophon, etc.
	"anc": true, // Announcer
	"anl": true, // Analyst
	"anm": true, // Animator
	"ann": true, // Annotator
	"ant": true, // Bibliographic antecedent
	"ape": true, // Appellee
	"apl": true, // Appellant
	"app": true, // Applicant
	"aqt": true, // Author in quotations or text abstracts
	"arc": true, // Architect
	"ard": true, // Artistic director
	"arr": true, // Arranger
	"art": true, // Artist
	"asg": true, // Assignee
	"asn": true, // Associated name
	"ato": true, // Autographer
	"att": true, // Attributed name
	"auc": true, // Auctioneer
	"aud": true, // Author of dialog
	"aue": true, // Audio engineer
	"aui": true, // Author of introduction, etc.
	"aup": true, // Audio producer
	"aus": true, // Screenwriter
	"aut": true, // Author
	"bdd": true, // Binding designer
	"bjd": true, // Bookjacket designer
	"bka": true, // Book artist
	"bkd": true, // Book designer
	"bkp": true, // Book producer
	"blw": true, // Blurb writer
	"bnd": true, // Binder
	"bpd": true, // Bookplate designer
	"brd": true, // Broadcaster
	"brl": true, // Braille embosser
	"bsl": true, // Bookseller
	"cad": true, // Casting director
	"cas": true, // Caster
	"ccp": true, // Conceptor
	"chr": true, // Choreographer
	"cli": true, // Client
	"cll": true, // Calligrapher
	"clr": true, // Colorist
	"clt": true, // Collotyper
	"cmm": true, // Commentator
	"cmp": true, // Composer
	"cmt": true, // Compositor
	"cnd": true, // Conductor
	"cng": true, // Cinematographer
	"cns": true, // Censor
	"coe": true, // Contestant-appellee
	"col": true, // Collector
	"com": true, // Compiler
	"con": true, // Conservator
	"cop": true, // Camera operator
	"cor": true, // Collection registrar
	"cos": true, // Contestant
	"cot": true, // Contestant-appellant
	"cou": true, // Court governed
	"cov": true, // Cover designer
	"cpc": true, // Copyright claimant
	"cpe": true, // Complainant-appellee
	"cph": true, // Copyright holder
	"cpl": true, // Complainant
	"cpt": true, // Complainant-appellant
	"cre": true, // Creator
	"crp": true, // Correspondent
	"crr": true, // Corrector
	"crt": true, // Court reporter
	"csl": true, // Consultant
	"csp": true, // Consultant to a project
	"cst": true, // Costume designer
	"ctb": true, // Contributor
	"cte": true, // Contestee-appellee
	"ctg": true, // Cartographer
	"ctr": true, // Contractor
	"cts": true, // Contestee
	"ctt": true, // Contestee-appellant
	"cur": true, // Curator
	"cwt": true, // Commentator for written text
	"dbd": true, // Dubbing director
	"dbp": true, // Distribution place
	"dfd": true, // Defendant
	"dfe": true, // Defendant-appellee
	"dft": true, // Defendant-appellant
	"dgc": true, // Degree committee member
	"dgg": true, // Degree granting institution
	"dgs": true, // Degree supervisor
	"dis": true, // Dissertant
	"djo": true, // Dj
	"dln": true, // Delineator
	"dnc": true, // Dancer
	"dnr": true, // Donor
	"dpc": true, // Depicted
	"dpt": true, // Depositor
	"drm": true, // Draftsman
	"drt": true, // Director
	"dsr": true, // Designer
	"dst": true, // Distributor
	"dtc": true, // Data contributor
	"dte": true, // Dedicatee
	"dtm": true, // Data manager
	"dto": true, // Dedicator
	"dub": true, // Dubious author
	"edc": true, // Editor of compilation
	"edd": true, // Editorial director
	"edm": true, // Editor of moving image work
	"edt": true, // Editor
	"egr": true, // Engraver
	"elg": true, // Electrician
	"elt": true, // Electrotyper
	"eng": true, // Engineer
	"enj": true, // Enacting jurisdiction
	"etr": true, // Etcher
	"evp": true, // Event place
	"exp": true, // Expert
	"fac": true, // Facsimilist
	"fds": true, // Film distributor
	"fld": true, // Field director
	"flm": true, // Film editor
	"fmd": true, // Film director
	"fmk": true, // Filmmaker
	"fmo": true, // Former owner
	"fmp": true, // Film producer
	"fnd": true, // Funder
	"fon": true, // Founder
	"fpy": true, // First party
	"frg": true, // Forger
	"gdv": true, // Game developer
	"gis": true, // Geographic information specialist
	"his": true, // Host institution
	"hnr": true, // Honoree
	"hst": true, // Host
	"ill": true, // Illustrator
	"ilu": true, // Illuminator
	"ink": true, // Inker
	"ins": true, // Inscriber
	"inv": true, // Inventor
	"isb": true, // Issuing body
	"itr": true, // Instrumentalist
	"ive": true, // Interviewee
	"ivr": true, // Interviewer
	"jud": true, // Judge
	"jug": true, // Jurisdiction governed
	"lbr": true, // Laboratory
	"lbt": true, // Librettist
	"ldr": true, // Laboratory director
	"led": true, // Lead
	"lee": true, // Libelee-appellee
	"lel": true, // Libelee
	"len": true, // Lender
	"let": true, // Libelee-appellant
	"lgd": true, // Lighting designer
	"lie": true, // Libelant-appellee
	"lil": true, // Libelant
	"lit": true, // Libelant-appellant
	"lsa": true, // Landscape architect
	"lse": true, // Licensee
	"lso": true, // Licensor
	"ltg": true, // Lithographer
	"ltr": true, // Letterer
	"lyr": true, // Lyricist
	"mcp": true, // Music copyist
	"mdc": true, // Metadata contact
	"med": true, // Medium
	"mfp": true, // Manufacture place
	"mfr": true, // Manufacturer
	"mka": true, // Makeup artist
	"mod": true, // Moderator
	"mon": true, // Monitor
	"mrb": true, // Marbler
	"mrk": true, // Markup editor
	"msd": true, // Musical director
	"mte": true, // Metal engraver
	"mtk": true, // Minute taker
	"mup": true, // Music programmer
	"mus": true, // Musician
	"mxe": true, // Mixing engineer
	"nan": true, // News anchor
	"nrt": true, // Narrator
	"onp": true, // Onscreen participant
	"opn": true, // Opponent
	"org": true, // Originator
	"orm": true, // Organizer
	"osp": true, // Onscreen presenter
	"oth": true, // Other
	"own": true, // Owner
	"pad": true, // Place of address
	"pan": true, // Panelist
	"pat": true, // Patron
	"pbd": true, // Publisher director
	"pbl": true, // Publisher
	"pdr": true, // Project director
	"pfr": true, // Proofreader
	"pht": true, // Photographer
	"plt": true, // Platemaker
	"pma": true, // Permitting agency
	"pmn": true, // Production manager
	"pnc": true, // Penciller
	"pop": true, // Printer of plates
	"ppm": true, // Papermaker
	"ppt": true, // Puppeteer
	"pra": true, // Praeses
	"prc": true, // Process contact
	"prd": true, // Production personnel
	"pre": true, // Presenter
	"prf": true, // Performer
	"prg": true, // Programmer
	"prm": true, // Printmaker
	"prn": true, // Production company
	"pro": true, // Producer
	"prp": true, // Production place
	"prs": true, // Production designer
	"prt": true, // Printer
	"prv": true, // Provider
	"pta": true, // Patent applicant
	"pte": true, // Plaintiff-appellee
	"ptf": true, // Plaintiff
	"pth": true, // Patent holder
	"ptt": true, // Plaintiff-appellant
	"pup": true, // Publication place
	"rap": true, // Rapporteur
	"rbr": true, // Rubricator
	"rcd": true, // Recordist
	"rce": true, // Recording engineer
	"rcp": true, // Addressee
	"rdd": true, // Radio director
	"red": true, // Redaktor
	"ren": true, // Renderer
	"res": true, // Researcher
	"rev": true, // Reviewer
	"rpc": true, // Radio producer
	"rps": true, // Repository
	"rpt": true, // Reporter
	"rpy": true, // Responsible party
	"rse": true, // Respondent-appellee
	"rsg": true, // Restager
	"rsp": true, // Respondent
	"rsr": true, // Restorationist
	"rst": true, // Respondent-appellant
	"rth": true, // Research team head
	"rtm": true, // Research team member
	"rxa": true, // Remix artist
	"sad": true, // Scientific advisor
	"sce": true, // Scenarist
	"scl": true, // Sculptor
	"scr": true, // Scribe
	"sde": true, // Sound engineer
	"sds": true, // Sound designer
	"sec": true, // Secretary
	"sfx": true, // Special effects provider
	"sgd": true, // Stage director
	"sgn": true, // Signer
	"sht": true, // Supporting host
	"sll": true, // Seller
	"sng": true, // Singer
	"spk": true, // Speaker
	"spn": true, // Sponsor
	"spy": true, // Second party
	"srv": true, // Surveyor
	"std": true, // Set designer
	"stg": true, // Setting
	"stl": true, // Storyteller
	"stm": true, // Stage manager
	"stn": true, // Standards body
	"str": true, // Stereotyper
	"swd": true, // Software developer
	"tad": true, // Technical advisor
	"tau": true, // Television writer
	"tcd": true, // Technical director
	"tch": true, // Teacher
	"ths": true, // Thesis advisor
	"tld": true, // Television director
	"tlg": true, // Television guest
	"tlh": true, // Television host
	"tlp": true, // Television producer
	"trc": true, // Transcriber
	"trl": true, // Translator
	"tyd": true, // Type designer
	"tyg": true, // Typographer
	"uvp": true, // University place
	"vac": true, // Voice actor
	"vdg": true, // Videographer
	"vfx": true, // Visual effects provider
	"voc": true, // Vocalist
	"wac": true, // Writer of added commentary
	"wal": true, // Writer of added lyrics
	"wam": true, // Writer of accompanying material
	"wat": true, // Writer of added text
	"waw": true, // Writer of afterword
	"wdc": true, // Woodcutter
	"wde": true, // Wood engraver
	"wfs": true, // Writer of film story
	"wft": true, // Writer of intertitles
	"wfw": true, // Writer of foreword
	"win": true, // Writer of introduction
	"wit": true, // Witness
	"wpr": true, // Writer of preface
	"wst": true, // Writer of supplementary textual content
	"wts": true, // Writer of television story
}
