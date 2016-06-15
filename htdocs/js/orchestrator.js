var graph;
function myGraph() {

  // Add and remove elements on the graph object
  this.addNode = function (id) {
    nodes.push({"id": id});
    update();
  };

  this.removeNode = function (id) {
    var i = 0;
    var n = findNode(id);
    while (i < links.length) {
      if ((links.i.source == n) || (links.i.target == n)) {
        links.splice(i, 1);
      }
      else i++;
    }
    nodes.splice(findNodeIndex(id), 1);
    update();
  };

  this.removeLink = function (source, target) {
    for (var i = 0; i < links.length; i++) {
      if (links[i].source.id == source && links[i].target.id == target) {
        links.splice(i, 1);
        break;
      }
    }
    update();
  };

  this.removeallLinks = function () {
    links.splice(0, links.length);
    update();
  };

  this.removeAllNodes = function () {
    nodes.splice(0, links.length);
    update();
  };

  this.addLink = function (source, target, value) {
    links.push({"source": findNode(source), "target": findNode(target), "value": value});
    update();
  };

  var findNode = function (id) {
    for (var i in nodes) {
      if (nodes.i.id === id) return nodes[i];
    }
  };

  var findNodeIndex = function (id) {
    for (var i = 0; i < nodes.length; i++) {
      if (nodes[i].id == id) {
        return i;
      }
    }
  };

  // set up the D3 visualisation in the specified element
  var w = 960,
    h = 450;

  var color = d3.scale.category10();

  var vis = d3.select("body")
    .append("svg:svg")
    .attr("width", w)
    .attr("height", h)
    .attr("id", "svg")
    .attr("pointer-events", "all")
    .attr("viewBox", "0 0 " + w + " " + h)
    .attr("perserveAspectRatio", "xMinYMid")
    .append('svg:g');

  var force = d3.layout.force();

  var nodes = force.nodes(),
    links = force.links();

  var update = function () {
    var link = vis.selectAll("line")
      .data(links, function (d) {
        return d.source.id + "-" + d.target.id;
      });

    link.enter().append("line")
      .attr("id", function (d) {
        return d.source.id + "-" + d.target.id;
      })
      .attr("stroke-width", function (d) {
        return d.value / 10;
      })
      .attr("class", "link");
    link.append("title")
      .text(function (d) {
        return d.value;
      });
    link.exit().remove();

    var node = vis.selectAll("g.node")
      .data(nodes, function (d) {
        return d.id;
      });

    var nodeEnter = node.enter().append("g")
      .attr("class", "node")
      .call(force.drag);

    nodeEnter.append("svg:circle")
      .attr("r", 12)
      .attr("id", function (d) {
        return "Node;" + d.id;
      })
      .attr("class", "nodeStrokeClass")
      .attr("fill", function(d) { return color(d.id); });

    nodeEnter.append("svg:text")
      .attr("class", "textClass")
      .attr("x", 14)
      .attr("y", ".31em")
      .text(function (d) {
        return d.id;
      });

    node.exit().remove();

    force.on("tick", function () {

      node.attr("transform", function (d) {
        return "translate(" + d.x + "," + d.y + ")";
      });

      link.attr("x1", function (d) {
        return d.source.x;
      })
        .attr("y1", function (d) {
          return d.source.y;
        })
        .attr("x2", function (d) {
          return d.target.x;
        })
        .attr("y2", function (d) {
          return d.target.y;
        });
    });

    // Restart the force layout.
    force
      .gravity(0.01)
      .charge(-400)
      .friction(0)
    //.linkDistance( function(d) { return d.value * 10 } )
      .linkDistance( 120)
      .size([w, h])
      .start();
  };


  // Make it all go
  update();
}

function drawGraph() {

  graph = new myGraph("#svgdiv");


  graph.addNode('Sophia');
  graph.addNode('Daniel');
  graph.addNode('Ryan');
  graph.addNode('Lila');
  graph.addNode('Suzie');
  graph.addNode('Riley');
  graph.addNode('Grace');
  graph.addNode('Dylan');
  graph.addNode('Mason');
  graph.addNode('Emma');
  graph.addNode('Alex');
  graph.addLink('Sophia', 'Daniel', '10');
  graph.addLink('Daniel', 'Ryan', '10');
  graph.addLink('Ryan', 'Lila', '10');
  graph.addLink('Lila', 'Suzie', '10');
  graph.addLink('Suzie', 'Riley', '10');
  keepNodesOnTop();

  // callback for the changes in the network
  var step = -1;
  function nextval()
  {
    step++;
    return 2000 + (1500*step); // initial time, wait time
  }

  setTimeout(function() {
    //graph.addLink('Alex', 'Sophia', '20');
    keepNodesOnTop();
  }, nextval());

  setTimeout(function() {
    //graph.addLink('Sophia', 'Daniel', '20');
    keepNodesOnTop();
  }, nextval());

  setTimeout(function() {
    //graph.addLink('Daniel', 'Alex', '20');
    keepNodesOnTop();
  }, nextval());

  setTimeout(function() {
    //graph.addLink('Suzie', 'Daniel', '30');
    keepNodesOnTop();
  }, nextval());

  setTimeout(function() {
    //graph.removeLink('Dylan', 'Mason');
    //graph.addLink('Dylan', 'Mason', '8');
    keepNodesOnTop();
  }, nextval());

  setTimeout(function() {
    //graph.removeLink('Dylan', 'Emma');
    //graph.addLink('Dylan', 'Emma', '8');
    keepNodesOnTop();
  }, nextval());

}

drawGraph();

// because of the way the network is created, nodes are created first, and links second,
// so the lines were on top of the nodes, this just reorders the DOM to put the svg:g on top
function keepNodesOnTop() {
  $(".nodeStrokeClass").each(function( index ) {
    var gnode = this.parentNode;
    gnode.parentNode.appendChild(gnode);
  });
}
function addNodes() {
  d3.select("svg")
    .remove();
  drawGraph();
}

var ws;
window.onload = function() {
  ws = new WebSocket('wss://' + window.location.host + '/orchestrator');
  ws.onopen = function() {
    var msg = {
      name: 'orchestrator',
      state: "initial",
      date: Date.now()
    };
    ws.send(JSON.stringify(msg));
  };

  // Write message on receive
  ws.onmessage = function(e) {
    console.log("Received:",e);
    //document.getElementById('output').innerHTML += "Received: " + e.data + "<br>";
    var obj = JSON.parse(e.data);
  };
  ws.onclose = function() {
    document.body.style.background = 'white';
  };
};

