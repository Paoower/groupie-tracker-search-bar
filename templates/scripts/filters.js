const data = {
    fromCreationYear: 1950,
    toCreationYear: 2024,
    fromAlbumYear: 1950,
    toAlbumYear: 2024,
    members: [],
    location: ""
}

console.log("zizi")

data.fromCreationYear = 1950
data.toCreationYear = 2024
data.fromAlbumYear = 1950
data.toAlbumYear = 2024

document.querySelectorAll("input[type=checkbox][name=members]").forEach(checkbox => {
    checkbox.addEventListener("change", e => {
        if (e.target.checked) {
            data.members.push(parseInt(e.target.value));
        } else {
            data.members = data.members.filter(value => value !== parseInt(e.target.value));
        }
        handleFilter();
    });
});

const handleFilter = async () => {
	reqOptions = {
		method: "POST",
		body: JSON.stringify(data)
	};
	const req = await fetch("/filters", reqOptions);
	const text = await req.text();
	let container = document.getElementById("artists-container").innerHTML
	if (text != container) {
		document.getElementById("artists-container").innerHTML = text;
	}
}

document.getElementById("filters-button").addEventListener("click", () => {
	document.getElementById("filters").classList.toggle("hidden")
})

document.querySelectorAll("input[type=range]").forEach(range => {
	range.addEventListener("input", async e => {
		if (e.target.id == "slider-1") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[1];
			if (parseInt(e.target.value) > parseInt(s.value)) {
				e.target.value = s.value;
			}
		}
		if (e.target.id == "slider-2") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[0];
			if (parseInt(e.target.value) < parseInt(s.value)) {
				e.target.value = s.value;
			}
		}

		if (e.target.id == "slider-3") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[1];
			if (parseInt(e.target.value) > parseInt(s.value)) {
				e.target.value = s.value;
			}
		}
		if (e.target.id == "slider-4") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[0];
			if (parseInt(e.target.value) < parseInt(s.value)) {
				e.target.value = s.value;
			}
		}

		data[e.target.name] = parseInt(e.target.value);
		labelsDiv = e.target.closest("div").querySelector("div");
		labelsDiv.querySelectorAll("label").forEach(label => {
			if (label.htmlFor == e.target.id) {
				label.innerHTML = e.target.value
			}
		})
		await handleFilter();
	})
})

document.querySelectorAll("select").forEach(select => {
	select.addEventListener("change", e => {
		if (e.target.name == "members") {
			data[e.target.name] = parseInt(e.target.value)
			handleFilter()
		} else if (e.target.name == "location") {
			data[e.target.name] = e.target.value
			handleFilter()
		} else {
			data[e.target.name] = e.target.value
			handleFilter()
		}
	})
})