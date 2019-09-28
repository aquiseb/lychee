let db = {
	post: [
		{
			id: "p1",
			title: "Title 1",
			content: "Lorem Ipsum 1",
			creatorId: "u2"
		},
		{
			id: "p2",
			title: "Title 2",
			content: "Lorem Ipsum 2",
			creatorId: "u3"
		},
		{
			id: "p3",
			title: "Title 3",
			content: "Lorem Ipsum 3",
			creatorId: "u2"
		},
		{
			id: "p4",
			title: "Title 4",
			content: "Lorem Ipsum 4",
			creatorId: "u1"
		},
		{
			id: "p5",
			title: "Title 5",
			content: "Lorem Ipsum 5",
			creatorId: "u1"
		},
		{
			id: "p6",
			title: "Title 6",
			content: "Lorem Ipsum 6",
			creatorId: "u2"
		},
		{
			id: "p7",
			title: "Title 7",
			content: "Lorem Ipsum 7",
			creatorId: "u1"
		},
		{
			id: "p8",
			title: "Title 8",
			content: "Lorem Ipsum 8",
			creatorId: "u2"
		}
	],
	user: [
		{
			id: "u1",
			firstname: "Bob",
			lastname: "Dylan"
		},
		{
			id: "u2",
			firstname: "Johnny",
			lastname: "Cash"
		},
		{
			id: "u3",
			firstname: "Elvis",
			lastname: "Presley"
		}
	]
};

// Imitate the way mongodb filters the results
const mongoStyleFilter = (dbQuery, document) => Object.keys(dbQuery).every(key => dbQuery[key] === document[key]);

// Some mongodb methods
const findOne = collectionName => (dbQuery = {}, callback = () => {}) => {
	const document = db[collectionName].find(document => mongoStyleFilter(dbQuery, document));
	callback(null, document);
	return document;
};

const find = collectionName => (dbQuery = {}, callback = () => {}) => {
	const document = db[collectionName].filter(document => mongoStyleFilter(dbQuery, document));
	callback(null, document);
	return document;
};

const updateOne = collectionName => (dbQuery = {}, dbUpdate = { $set: {} }, callback = () => {}) => {
	const indexToUpdate = db[collectionName].findIndex(document => mongoStyleFilter(dbQuery, document));
	const collection = db[collectionName];
	if (indexToUpdate > -1) {
		collection[indexToUpdate] = {
			...collection[indexToUpdate],
			...dbUpdate.$set
		};
	}

	callback(null, collection[indexToUpdate]);
	return collection[indexToUpdate];
};

// Create a collectionNames object with each method
let collectionNames = ["post", "user"].reduce((acc, curr) => {
	acc[curr] = {
		find: find(curr),
		findOne: findOne(curr),
		updateOne: updateOne(curr)
	};
	return acc;
}, {});

export default collectionNames;
